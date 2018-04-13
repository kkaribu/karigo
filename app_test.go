package karigo

import (
	"sort"
	"testing"
	"time"

	"github.com/kkaribu/jsonapi"
	"github.com/kkaribu/tchek"
)

func TestIntegrityCheck(t *testing.T) {
	app := NewApp(nil)

	app.RegisterType(&ValidType1{})
	app.RegisterType(&ValidType2{})
	app.RegisterType(&ValidType3{})
	app.RegisterType(&InvalidType1{})

	errs := app.Check()
	expectations := []string{
		// 1
		"jsonapi: the inverse of relationship rel1 of type invalidtype1 does not exist",
		// 2
		"jsonapi: the inverse of relationship rel2 of type invalidtype1 does not exist",
		// 3
		"jsonapi: the target type of relationship rel2 of type invalidtype1 does not exist",
	}

	errStrs := []string{}
	for _, e := range errs {
		errStrs = append(errStrs, e.Error())
	}
	sort.Strings(errStrs)

	// 0
	tchek.AreEqual(t, 0, len(errStrs), len(expectations))

	if len(errStrs) == len(expectations) {
		for i := 0; i < len(errStrs); i++ {
			tchek.AreEqual(t, i+1, expectations[i], errStrs[i])
		}
	}

}

func TestInstanceDistributor(t *testing.T) {

}

func TestInfo(t *testing.T) {
	app := NewApp(nil)

	app.Info("message")

	app.Info("message %s", "with string argument")
}

type ValidType1 struct {
	ID string `json:"id" api:"validtype1"`

	Attr1 string     `json:"attr1" api:"attr"`
	Attr2 uint       `json:"attr2" api:"attr"`
	Attr3 bool       `json:"attr3" api:"attr"`
	Attr4 time.Time  `json:"attr4" api:"attr"`
	Attr5 *string    `json:"attr5" api:"attr"`
	Attr6 *uint      `json:"attr6" api:"attr"`
	Attr7 *bool      `json:"attr7" api:"attr"`
	Attr8 *time.Time `json:"attr8" api:"attr"`

	Rel1 string   `json:"rel1" api:"rel,validtype2"`
	Rel2 []string `json:"rel2" api:"rel,validtype3"`
	Rel3 string   `json:"rel3" api:"rel,validtype2,rel1"`
	Rel4 []string `json:"rel4" api:"rel,validtype3,rel1"`
}

type ValidType2 struct {
	ID string `json:"id" api:"validtype2"`

	Attr1 string `json:"attr1" api:"attr"`

	Rel1 string `json:"rel1" api:"rel,validtype1,rel3"`
}

type ValidType3 struct {
	ID string `json:"id" api:"validtype3"`

	Attr1 string `json:"attr1" api:"attr"`

	Rel1 []string `json:"rel1" api:"rel,validtype1,rel4"`
}

type InvalidType1 struct {
	ID string `json:"id" api:"invalidtype1"`

	Attr1 string `json:"attr1" api:"attr"`

	Rel1 []string `json:"rel1" api:"rel,validtype1,rel99"`
	Rel2 []string `json:"rel2" api:"rel,validtype99,rel1"`
}

// NewMockApp ...
func NewMockApp() *App {
	app := NewApp(nil)

	app.Registry = jsonapi.NewMockRegistry()

	app.Get("/mocktypes3", KernelGetCollection, GatePublic)

	return app
}
