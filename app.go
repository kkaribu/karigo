package karigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/kkaribu/jsonapi"
	"gopkg.in/rs/cors.v1"
	"gopkg.in/urfave/cli.v1"
)

// NewApp creates and returns an App object.
func NewApp(store Store) *App {
	app := &App{
		Store: store,
		CLI:   cli.NewApp(),

		Registry: jsonapi.NewRegistry(),

		Hooks:   map[string]func(){},
		Kernels: map[string]Kernel{},
		Gates:   map[string][]Gate{},
	}

	return app
}

// An App represents the instance of an application that listens and responds
// to HTTP requests.
type App struct {
	sync.Mutex

	Config Config
	Store  Store    `json:"-"`
	CLI    *cli.App `json:"-"`

	*jsonapi.Registry

	Hooks   map[string]func() `json:"-"`
	Kernels map[string]Kernel `json:"-"`
	Gates   map[string][]Gate `json:"-"`
}

// Info ...
func (a *App) Info(msg string, args ...interface{}) {
	log.Printf(msg, args...)
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

// Merge ...
func (a *App) Merge(na *App) {
	for n, t := range na.Types {
		a.Types[n] = t
	}

	for n, f := range na.Hooks {
		a.Hooks[n] = f
	}

	for n, k := range na.Kernels {
		a.Kernels[n] = k
	}

	for n, g := range na.Gates {
		a.Gates[n] = g
	}
}

// RunCLI ...
func (a *App) RunCLI() {
	a.CLI.Name = a.Config.Name

	a.CLI.Metadata = map[string]interface{}{}
	a.CLI.Metadata["app"] = a

	a.AddCmd(
		deleteCmd(),
		runCmd(),
		schemaCmd(),
		syncCmd(),
	)

	err := a.CLI.Run(os.Args)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Println()
}

// AddHook ...
func (a *App) AddHook(pos string, f func()) {
	a.Hooks[pos] = f
}

// RunHook ...
func (a *App) RunHook(pos string) {
	if f, ok := a.Hooks[pos]; ok {
		f()
	}
}

// AddCmd ...
func (a *App) AddCmd(cmd ...cli.Command) {
	a.CLI.Commands = append(a.CLI.Commands, cmd...)
}

// Run ...
func (a *App) Run() error {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	})

	handler := c.Handler(a)

	http.ListenAndServe(fmt.Sprintf(":%d", a.Config.Port), handler)

	return nil
}

// Schema ...
func (a *App) Schema() string {
	info, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}

	return string(info)
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
				jaerr = jsonapi.NewErrInternal()
			case jsonapi.Error:
				ctx.AddToLog(fmt.Sprintf("JSONAPI error: %s\n", e.Error()))
				jaerr = e
			case error:
				ctx.AddToLog(fmt.Sprintf("Error: %s\n", e))
				jaerr = jsonapi.NewErrInternal()
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

	ctx.AddToLog("Context initialized.")

	// Parse URL
	r.URL.Host = r.Host
	if r.TLS != nil {
		r.URL.Scheme = "https"
	} else {
		r.URL.Scheme = "http"
	}
	// The host and scheme are set there because they might not be set by
	// default if the request was made with a relative path.
	url, err := jsonapi.ParseURL(a.Registry, r.URL)
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
			panic(jsonapi.NewErrInternal())
		}
		ctx.In, err = jsonapi.Unmarshal(body, ctx.URL, ctx.App.Registry)
		if err != nil {
			panic(jsonapi.NewErrBadRequest())
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
		panic(jsonapi.NewErrInternal())
	}
	ctx.Tx = tx

	// Execute kernel
	a.executeKernel(ctx)

	ctx.AddToLog("Kernel executed.")

	// Check Document
	body, err = jsonapi.Marshal(ctx.Out, ctx.URL)
	if err != nil {
		panic(jsonapi.NewErrInternal())
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
