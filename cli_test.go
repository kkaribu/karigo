package karigo

import (
	"bytes"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/kkaribu/jsonapi"
	"github.com/kkaribu/tchek"
	"gopkg.in/urfave/cli.v1"
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
	tchek.ErrorExpected(t, 0, false, err)

	err = cliApp.Run([]string{"karigo", "help"})
	tchek.ErrorExpected(t, 1, false, err)

	err = cliApp.Run([]string{"karigo", "--config=tests/invalid.json"})
	tchek.ErrorExpected(t, 2, true, err)

	err = cliApp.Run([]string{"karigo", "schema"})
	tchek.ErrorExpected(t, 3, true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config"})
	tchek.ErrorExpected(t, 4, true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config="})
	tchek.ErrorExpected(t, 5, true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config=tests/invalid.json"})
	tchek.ErrorExpected(t, 6, true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/invalid.json"})
	tchek.ErrorExpected(t, 7, true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, 8, false, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config=tests/config_valid2.json"})
	tchek.ErrorExpected(t, 9, false, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_valid1.json"})
	tchek.ErrorExpected(t, 10, false, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_invalid1.json"})
	tchek.ErrorExpected(t, 11, true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_invalid2.json"})
	tchek.ErrorExpected(t, 12, true, err)

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_invalid3.json"})
	tchek.ErrorExpected(t, 13, true, err)

	cliApp = makeCLIAppWithMockStore(&MockStore{
		OpenFunc: func(driver, host, db, user, pw string, opts map[string]string) error {
			return errors.New("error")
		},
	})

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_valid1.json"})
	tchek.ErrorExpected(t, 14, true, err)

	invalidApp := NewMockApp()
	invalidApp.RegisterType(InvalidType1{})
	cliApp.Metadata["app"] = invalidApp

	err = cliApp.Run([]string{"karigo", "schema", "--config", "tests/config_valid1.json"})
	tchek.ErrorExpected(t, 15, true, err)
}

func TestCheckCmd(t *testing.T) {
	t.Parallel()

	cliApp := makeCLIAppWithMockStore(&MockStore{
		SyncDatabaseFunc: func(tx Tx, reg *jsonapi.Registry, verbose, apply bool) error {
			return nil
		},
	})

	err := cliApp.Run([]string{"karigo", "check", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, 0, false, err)

	cliApp = makeCLIAppWithMockStore(&MockStore{
		SyncDatabaseFunc: func(tx Tx, reg *jsonapi.Registry, verbose, apply bool) error {
			return errors.New("error")
		},
	})

	err = cliApp.Run([]string{"karigo", "check", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, 1, true, err)

	err = cliApp.Run([]string{"karigo", "check", "--config", "tests/invalid.json"})
	tchek.ErrorExpected(t, 2, true, err)
}

func TestDrainCmd(t *testing.T) {
	t.Parallel()

	cliApp := makeCLIAppWithMockStore(&MockStore{
		DrainDatabaseFunc: func(tx Tx) error {
			return nil
		},
	})

	err := cliApp.Run([]string{"karigo", "drain", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, 0, false, err)

	cliApp = makeCLIAppWithMockStore(&MockStore{
		DrainDatabaseFunc: func(tx Tx) error {
			return errors.New("error")
		},
	})

	err = cliApp.Run([]string{"karigo", "drain", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, 1, true, err)

	err = cliApp.Run([]string{"karigo", "drain", "--config", "tests/invalid.json"})
	tchek.ErrorExpected(t, 2, true, err)
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
	tchek.ErrorExpected(t, 0, false, err)
	<-c

	err = cliApp.Run([]string{"karigo", "run", "--config", "tests/invalid.json"})
	tchek.ErrorExpected(t, 1, true, err)
}

func TestSchemaCmd(t *testing.T) {
	t.Parallel()

	cliApp := makeCLIAppWithMockStore(&MockStore{})

	err := cliApp.Run([]string{"karigo", "schema", "--config=tests/config_valid1.json"})
	tchek.ErrorExpected(t, 0, false, err)
}
