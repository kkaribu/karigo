package karigo

import (
	"fmt"
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
			fmt.Println("Draining database...")
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
			fmt.Println(app.Schema())

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

			if app.Config.Debug {
				fmt.Println("Debug: on")
			} else {
				fmt.Println("Debug: off")
			}
			fmt.Printf("Now listening on %d...\n", app.Config.Port)
			fmt.Println()

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

	// Check app
	errs := app.Check()
	if len(errs) > 0 {
		return app, errs[0]
	}

	return app, nil
}

// TerminateCmd ...
func TerminateCmd() {
	fmt.Println()
}
