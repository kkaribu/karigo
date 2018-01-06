package karigo

import (
	"io/ioutil"

	"gopkg.in/urfave/cli.v1"
)

func deleteCmd() cli.Command {
	return cli.Command{
		Name:    "delete",
		Aliases: []string{},
		Usage:   "delete instance",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			app, err := PrepareCmd(c)
			if err != nil {
				return err
			}

			// Drain database
			app.Info("Draining database...")
			err = app.Store.DrainDatabase(nil)
			if err != nil {
				return err
			}

			TerminateCmd()

			return nil
		},
	}
}

func syncCmd() cli.Command {
	return cli.Command{
		Name:    "sync",
		Aliases: []string{},
		Usage:   "update the database's schema to match the app",
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

			// Sync database
			err = app.Store.SyncDatabase(nil, app.Registry, true, c.Bool("apply"))
			if err != nil {
				return err
			}

			TerminateCmd()

			return nil
		},
	}
}

func schemaCmd() cli.Command {
	return cli.Command{
		Name:    "schema",
		Aliases: []string{},
		Usage:   "show the app's schema",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			app, err := PrepareCmd(c)
			if err != nil {
				return err
			}

			// Info
			app.Info("\n" + app.Schema())

			TerminateCmd()

			return nil
		},
	}
}

func runCmd() cli.Command {
	return cli.Command{
		Name:    "run",
		Aliases: []string{},
		Usage:   "run instance",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			app, err := PrepareCmd(c)
			if err != nil {
				return err
			}

			app.RunHook("before-run")

			app.Info("Now listening on %d...\n\n", app.Config.Port)

			app.Run()

			TerminateCmd()

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
	// defer app.Store.Close() // TODO Where do we close it?
	if err != nil {
		return app, err
	}
	app.Info("URL: %s", app.Store.URL())
	app.Info("Connection to database established.")

	return app, nil
}

// TerminateCmd ...
func TerminateCmd() {
}
