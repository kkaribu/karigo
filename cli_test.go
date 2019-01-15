package karigo

import (
	"bytes"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/kkaribu/jsonapi"
	"github.com/kkaribu/tchek"
	"github.com/urfave/cli"
)

func makeCLIAppWithMockStore(ms *MockStore) *cli.App {
	if ms.OpenFunc == nil {
		ms.OpenFunc = func(driver, host, db, user, pw string, opts map[string]string) error {
			return nil
		}
	}

	if ms.URLFunc == nil {
		ms.URLFunc = func() string {
			return ""
		}
	}

	app := NewMockApp()
	app.Store = ms

	cliApp := cli.NewApp()
	cliApp.Metadata = map[string]interface{}{
		"app": app,
	}

	cliApp.Commands = []cli.Command{
		checkCmd(),
		drainCmd(),
		schemaCmd(),
		runCmd(),
	}

	cliApp.Writer = &bytes.Buffer{}

	return cliApp
}

func TestCLI(t *testing.T) {
	t.Parallel()

	cliApp := makeCLIAppWithMockStore(&MockStore{
		OpenFunc: func(driver, host, db, user, pw string, opts map[string]string) error {
			return nil
		},
		URLFunc: func() string {
			return ""
		},
	})

	err := cliApp.Run([]string{"karigo"})
	tchek.ErrorExpected(t, "karigo", false, err)

	err = cliApp.Run([]string{"karigo", "help"})
	tchek.ErrorExpected(t, "karigo help", false, err)

	err = cliApp.Run([]string{"karigo", "--config=tests/invalid.json"})
	tchek.ErrorExpected(t, "karigo invalid config", true, err)

	err = cliApp.Run([]string{"karigo", "schema"})
	tchek.ErrorExpected(t, "karigo schema", true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config"})
	tchek.ErrorExpected(t, "karigo schema config", true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config="})
	tchek.ErrorExpected(t, "karigo schema config 2", true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config=tests/invalid.json"})
	tchek.ErrorExpected(t, "karigo schema non-existent config", true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/invalid.json"})
	tchek.ErrorExpected(t, "karigo schema non-existent config 2", true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, "karigo schema config", false, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config=tests/config_valid2.json"})
	tchek.ErrorExpected(t, "karigo schema config 2", false, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_valid1.json"})
	tchek.ErrorExpected(t, "karigo schema config 3", false, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_invalid1.json"})
	tchek.ErrorExpected(t, "karigo schema invald config", true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_invalid2.json"})
	tchek.ErrorExpected(t, "karigo schema invald config 2", true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_invalid3.json"})
	tchek.ErrorExpected(t, "karigo schema invald config 3", true, err)

	cliApp = makeCLIAppWithMockStore(&MockStore{
		OpenFunc: func(driver, host, db, user, pw string, opts map[string]string) error {
			return errors.New("error")
		},
	})

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_valid1.json"})
	tchek.ErrorExpected(t, "karigo schema connection error", true, err)

	invalidApp := NewMockApp()
	invalidApp.RegisterType(InvalidType1{})
	cliApp.Metadata["app"] = invalidApp

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_valid1.json"})
	tchek.ErrorExpected(t, "karigo schema invalid app", true, err)
}

func TestCheckCmd(t *testing.T) {
	t.Parallel()

	cliApp := makeCLIAppWithMockStore(&MockStore{
		SyncDatabaseFunc: func(tx Tx, reg *jsonapi.Registry, verbose, apply bool) error {
			return nil
		},
	})

	err := cliApp.Run([]string{"karigo", "check", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, "karigo check successful sync", false, err)

	cliApp = makeCLIAppWithMockStore(&MockStore{
		SyncDatabaseFunc: func(tx Tx, reg *jsonapi.Registry, verbose, apply bool) error {
			return errors.New("error")
		},
	})

	err = cliApp.Run([]string{"karigo", "check", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, "karigo check sync error", true, err)

	err = cliApp.Run([]string{"karigo", "check", "--config", "tests/invalid.json"})
	tchek.ErrorExpected(t, "karigo check non-existent config", true, err)
}

func TestDrainCmd(t *testing.T) {
	t.Parallel()

	cliApp := makeCLIAppWithMockStore(&MockStore{
		DrainDatabaseFunc: func(tx Tx) error {
			return nil
		},
	})

	err := cliApp.Run([]string{"karigo", "drain", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, "karigo drain successful drain", false, err)

	cliApp = makeCLIAppWithMockStore(&MockStore{
		DrainDatabaseFunc: func(tx Tx) error {
			return errors.New("error")
		},
	})

	err = cliApp.Run([]string{"karigo", "drain", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, "karigo drain drain error", true, err)

	err = cliApp.Run([]string{"karigo", "drain", "--config", "tests/invalid.json"})
	tchek.ErrorExpected(t, "karigo drain non-existent config", true, err)
}

func TestRunCmd(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	cliApp := makeCLIAppWithMockStore(&MockStore{})

	c := make(chan bool, 1)
	go func() {
		time.Sleep(time.Millisecond * 20)
		cliApp.Metadata["app"].(*App).Shutdown()
		c <- true
	}()
	err := cliApp.Run([]string{"karigo", "run", "--config=tests/config_valid2.json"})
	tchek.ErrorExpected(t, "karigo run successful", false, err)
	<-c

	err = cliApp.Run([]string{"karigo", "run", "--config", "tests/invalid.json"})
	tchek.ErrorExpected(t, "karigo run non-existent config", true, err)
}

func TestSchemaCmd(t *testing.T) {
	t.Parallel()

	cliApp := makeCLIAppWithMockStore(&MockStore{})

	err := cliApp.Run([]string{"karigo", "schema", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, "karigo schema successful", false, err)
}
