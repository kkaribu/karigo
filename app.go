package karigo

import (
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

	Name  string   `json:"name"`
	Port  uint16   `json:"port"`
	Debug bool     `json:"debug"`
	Store Store    `json:"-"`
	CLI   *cli.App `json:"-"`

	*jsonapi.Registry

	Hooks   map[string]func() `json:"-"`
	Kernels map[string]Kernel `json:"-"`
	Gates   map[string][]Gate `json:"-"`
}

// Info ...
func (a *App) Info(msg string) {
	log.Printf("%s", msg)
}

// ReadConfig ...
func (a *App) ReadConfig(data []byte) error {
	config := struct {
		Name  string
		Port  uint16
		Debug bool
		Store struct {
			Name     string
			Address  string
			Port     uint16
			User     string
			Password string
		}
	}{}

	err := json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	if config.Name == "" {
		return fmt.Errorf("karigo: missing app name in configuration file")
	}

	if config.Store.Name == "" {
		return fmt.Errorf("karigo: missing store name in configuration file")
	}

	if config.Store.Address == "" {
		return fmt.Errorf("karigo: missing store address in configuration file")
	}

	if config.Store.User == "" {
		return fmt.Errorf("karigo: missing store user in configuration file")
	}

	a.Name = config.Name
	a.Port = config.Port
	a.Debug = config.Debug

	// Connect to database
	a.Info("Connecting to database...")
	err = a.Store.Open(config.Store.User, config.Store.Password, config.Store.Address, a.Name)
	// defer app.Store.Close() // TODO Where do we close it?
	if err != nil {
		return err
	}

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
	a.CLI.Name = a.Name

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

	http.ListenAndServe(fmt.Sprintf(":%d", a.Port), handler)

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
			switch err := rec.(type) {
			case string:
				ctx.AddToLog(fmt.Sprintf("String error: %s\n", err))
				jaerr = jsonapi.NewErrInternal()
			case jsonapi.Error:
				ctx.AddToLog(fmt.Sprintf("JSONAPI error: %s\n", err.Error()))
				jaerr = err
			case error:
				ctx.AddToLog(fmt.Sprintf("Error: %s\n", err))
				jaerr = jsonapi.NewErrInternal()
			}

			w.WriteHeader(jaerr.Status)

			var body []byte
			// var err error

			// body, err = jsonapi.Marshal(jaerr, &jsonapi.URL{}, nil)
			// if err != nil {
			body = []byte("{\"errors\":{\"title\":\"Epic Fail\"}}")
			// }

			_, _ = w.Write(body)

			fmt.Println()
		}

		// Print log
		if a.Debug {
			ctx.SaveLog()
		}
	}()

	// Initialize context
	ctx.Store = a.Store
	ctx.Method = r.Method
	ctx.Body, _ = ioutil.ReadAll(r.Body)
	ctx.Doc = jsonapi.NewDocument()

	ctx.AddToLog("Context initialized.")

	// Parse URL
	ctx.URL, _ = jsonapi.ParseURL(a.Registry, r.URL)
	ctx.Doc.Meta["interpreted-url"] = ctx.URL.URLNormalized

	ctx.AddToLog("URL parsed.")

	// ctx.Doc.Fields = ctx.URL.Params.Fields
	ctx.Doc.RelData = ctx.URL.Params.RelData

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
	var body []byte
	body, err = ctx.Doc.MarshalJSON()
	if err != nil {
		panic(jsonapi.NewErrInternal())
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
