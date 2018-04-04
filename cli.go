package karigo

import (
	"io/ioutil"

	"gopkg.in/urfave/cli.v1"
)

func drainCmd() cli.Command {
	return cli.Command{
		Name:    "drain",
		Aliases: []string{},
		Usage:   "Empties the store (including the tables if necessary), but keeps the store itself.",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			app, err := PrepareCmd(c)
			if err != nil {
				return err
			}

			// Drain store
			app.Info("Draining store...")
			err = app.Store.DrainDatabase(nil)
			if err != nil {
				return err
			}
			app.Info("Store is now drained.")

			TerminateCmd(app)

			return nil
		},
	}
}

func checkCmd() cli.Command {
	return cli.Command{
		Name:    "check",
		Aliases: []string{},
		Usage:   "Displays the differences between the app schema and the store schema.",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "apply",
				Usage: "apply the updates",
			},
		},
		Action: func(c *cli.Context) error {
			app, err := PrepareCmd(c)
			if err != nil {
				return err
			}

			// Sync store
			app.Info("Syncing store...")
			err = app.Store.SyncDatabase(nil, app.Registry, true, c.Bool("apply"))
			if err != nil {
				return err
			}
			app.Info("Store schema is synced.")

			TerminateCmd(app)

			return nil
		},
	}
}

func schemaCmd() cli.Command {
	return cli.Command{
		Name:    "schema",
		Aliases: []string{},
		Usage:   "Displays the app schema.",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			app, err := PrepareCmd(c)
			if err != nil {
				return err
			}

			// Info
			app.Info("\n" + app.Schema())

			TerminateCmd(app)

			return nil
		},
	}
}

func runCmd() cli.Command {
	return cli.Command{
		Name:    "run",
		Aliases: []string{},
		Usage:   "Instantiates the app and starts serving requests.",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			app, err := PrepareCmd(c)
			if err != nil {
				return err
			}

			app.RunHook("before-run")

			app.Info("Now listening on %d...\n\n", app.Config.Port)

			app.Run()

			TerminateCmd(app)

			return nil
		},
	}
}

// PrepareCmd ...
func PrepareCmd(c *cli.Context) (*App, error) {
	app := c.App.Metadata["app"].(*App)

	// Configuration
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return app, err
	}

	err = app.ReadConfig(data)
	if err != nil {
		return app, err
	}

	if app.Config.Debug {
		app.Info("Debug is ON.")
	} else {
		app.Info("Debug is OFF.")
	}

	// Check app
	errs := app.Check()
	if len(errs) > 0 {
		return app, errs[0]
	}

	// Connect to database
	app.Info("Connecting to database...")
	err = app.Store.Open(
		app.Config.Store.Driver,
		app.Config.Store.Host,
		app.Config.Store.Database,
		app.Config.Store.User,
		app.Config.Store.Password,
		app.Config.Store.Options,
	)
	app.Info("URL: %s", app.Store.URL())
	if err != nil {
		return app, err
	}
	app.Info("Connection to database established.")

	return app, nil
}

// TerminateCmd ...
func TerminateCmd(app *App) error {
	app.Info("Closing database connection...")
	app.Store.Close()
	app.Info("Connection closed.")
	return nil
}
