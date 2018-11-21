package karigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/kkaribu/jsonapi"
	"github.com/rs/cors"
	"github.com/urfave/cli"
)

// NewApp creates and returns an App object.
func NewApp(store Store) *App {
	app := &App{
		Log:    &Log{},
		Store:  store,
		CLI:    cli.NewApp(),
		Server: http.Server{},

		Registry: jsonapi.NewRegistry(),

		Actions: map[string]Action{},
		Kernels: map[string]Kernel{},
		Gates:   map[string][]Gate{},
	}

	app.Config.Store.Options = map[string]string{}

	return app
}

// An App represents the instance of an application that listens and responds
// to HTTP requests.
type App struct {
	sync.Mutex

	Config Config
	Log    *Log        `json:"-"`
	Store  Store       `json:"-"`
	CLI    *cli.App    `json:"-"`
	Server http.Server `json:"-"`

	*jsonapi.Registry

	Actions map[string]Action `json:"-"`
	Kernels map[string]Kernel `json:"-"`
	Gates   map[string][]Gate `json:"-"`
}

// Info ...
func (a *App) Info(msg string, args ...interface{}) {
	if a.Config.Info {
		log.Printf(msg, args...)
	}
}

// Debug ...
func (a *App) Debug(msg string, args ...interface{}) {
	if a.Config.Debug {
		log.Printf(msg, args...)
	}
}

// ReadConfig ...
func (a *App) ReadConfig(data []byte) error {
	config := Config{}

	err := json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	if config.Store.Driver == "" {
		return fmt.Errorf("karigo: missing store driver in configuration file")
	}

	if config.Store.Host == "" {
		return fmt.Errorf("karigo: missing store host in configuration file")
	}

	a.Config = config

	return nil
}

// RunCLI ...
func (a *App) RunCLI() {
	a.CLI.Name = a.Config.Name

	a.CLI.Metadata = map[string]interface{}{}
	a.CLI.Metadata["app"] = a

	a.CLI.Commands = append(a.CLI.Commands, drainCmd())
	a.CLI.Commands = append(a.CLI.Commands, runCmd())
	a.CLI.Commands = append(a.CLI.Commands, schemaCmd())
	a.CLI.Commands = append(a.CLI.Commands, checkCmd())

	err := a.CLI.Run(os.Args)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Println()
}

// Run ...
func (a *App) Run() error {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	})

	a.Server.Addr = ":" + strconv.FormatUint(uint64(a.Config.Port), 10)
	a.Server.Handler = c.Handler(a)

	err := a.Server.ListenAndServe()

	return err
}

// Shutdown ...
func (a *App) Shutdown() error {
	return a.Server.Shutdown(nil)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewCtx(a, WrapResponseWriter(w), r)
	ctx.AddToLog("Request starting...")

	// Handle panic
	defer func() {
		if rec := recover(); rec != nil {
			ctx.AddToLog("Panic!")

			var jaerr jsonapi.Error
			switch e := rec.(type) {
			case string:
				ctx.AddToLog(fmt.Sprintf("String error: %s\n", e))
				jaerr = jsonapi.NewErrInternalServerError()
			case jsonapi.Error:
				ctx.AddToLog(fmt.Sprintf("JSONAPI error: %s\n", e.Error()))
				jaerr = e
			case error:
				ctx.AddToLog(fmt.Sprintf("Error: %s\n", e))
				jaerr = jsonapi.NewErrInternalServerError()
			}

			ctx.Out = jsonapi.NewDocument()
			ctx.Out.Data = []jsonapi.Error{jaerr}

			var body []byte
			var err error
			body, err = jsonapi.Marshal(ctx.Out, ctx.URL)
			if err != nil {
				body = []byte(`{"errors":{"title":"Epic Fail"}}`)
			}

			w.WriteHeader(jaerr.Status)
			_, _ = w.Write(body)

			fmt.Println()
		}

		// Print log
		if a.Config.Debug {
			ctx.SaveLog()
		}
	}()

	// Initialize context
	ctx.Method = r.Method
	ctx.Out = jsonapi.NewDocument()
	ctx.Out.PrePath = a.Config.PrePath

	ctx.AddToLog("Context initialized.")

	// Parse URL
	url, err := jsonapi.ParseRawURL(a.Registry, r.URL.String())
	if err != nil {
		panic(err)
	}
	ctx.URL = url

	ctx.AddToLog("URL parsed.")

	// Parse body
	var body []byte
	if ctx.Method == "POST" || ctx.Method == "PATCH" {
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			panic(jsonapi.NewErrInternalServerError())
		}
		ctx.In, err = jsonapi.Unmarshal(body, ctx.URL, ctx.App.Registry)
		if err != nil {
			panic(jsonapi.NewErrBadRequest("Bad body", "The body is invalid."))
		}
	}

	// ctx.Out.Fields = ctx.URL.Params.Fields/
	ctx.Out.RelData = ctx.URL.Params.RelData

	// Defaults
	if ctx.URL.Params.PageNumber <= 0 {
		ctx.URL.Params.PageNumber = 1
	}
	if ctx.URL.Params.PageSize <= 0 {
		ctx.URL.Params.PageSize = 1000
	}

	// Parse JWT
	a.parseJWT(ctx, r)

	ctx.AddToLog("JWT parsed.")

	// Check gates
	a.checkGates(ctx)

	ctx.AddToLog("Gates checked.")

	// jctx, _ := json.MarshalIndent(ctx, "", "  ")
	// fmt.Printf("CONTEXT\n\n%s\n", jctx)

	// Begin transaction
	tx, err := a.Store.Begin()
	if err != nil {
		panic(jsonapi.NewErrInternalServerError())
	}
	ctx.Tx = tx

	// Execute kernel
	a.executeKernel(ctx)

	ctx.AddToLog("Kernel executed.")

	// Check Document
	body, err = jsonapi.Marshal(ctx.Out, ctx.URL)
	if err != nil {
		panic(jsonapi.NewErrInternalServerError())
	}
	if a.Config.Minimize {
		buf := &bytes.Buffer{}
		err = json.Indent(buf, body, "", "\t")
		if err != nil {
			panic(err)
		}
		body = buf.Bytes()
	}
	ctx.AddToLog("Document checked.")

	// Commit transaction
	err = ctx.Tx.Commit()
	if err != nil {
		panic("could not commit transaction")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)

	ctx.AddToLog("Response sent.")
}

// Get adds a kernel for handling GET requests.
func (a *App) Get(route string, k Kernel, g ...Gate) {
	a.Lock()
	defer a.Unlock()

	a.Kernels["GET "+route] = k

	if g != nil {
		a.Gates["GET "+route] = g
	}
}

// Post adds a kernel for handling POST requests.
func (a *App) Post(route string, k Kernel, g ...Gate) {
	a.Lock()
	defer a.Unlock()

	a.Kernels["POST "+route] = k

	if g != nil {
		a.Gates["POST "+route] = g
	}
}

// Put adds a kernel for handling PUT requests.
func (a *App) Put(route string, k Kernel, g ...Gate) {
	a.Lock()
	defer a.Unlock()

	a.Kernels["PUT "+route] = k

	if g != nil {
		a.Gates["PUT "+route] = g
	}
}

// Delete adds a kernel for handling DELETE requestsn.
func (a *App) Delete(route string, k Kernel, g ...Gate) {
	a.Lock()
	defer a.Unlock()

	a.Kernels["DELETE "+route] = k

	if g != nil {
		a.Gates["DELETE "+route] = g
	}
}
