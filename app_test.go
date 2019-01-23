package karigo

import (
	"sort"
	"testing"
	"time"

	"github.com/kkaribu/jsonapi"
	"github.com/kkaribu/tchek"
)

func TestIntegrityCheck(t *testing.T) {
	// app := NewApp(nil)

	reg := jsonapi.NewRegistry()

	reg.RegisterType(&ValidType1{})
	reg.RegisterType(&ValidType2{})
	reg.RegisterType(&ValidType3{})
	reg.RegisterType(&InvalidType1{})

	errs := reg.Check()
	tests := []struct {
		name string
		msg  string
	}{
		{
			name: "inverse rel does not exist",
			msg:  "jsonapi: the inverse of relationship rel1 of type invalidtype1 does not exist",
		}, {
			name: "inverse rel does not exist 2",
			msg:  "jsonapi: the inverse of relationship rel2 of type invalidtype1 does not exist",
		}, {
			name: "target type of relationship does not exist",
			msg:  "jsonapi: the target type of relationship rel2 of type invalidtype1 does not exist",
		},
	}

	errStrs := []string{}
	for _, e := range errs {
		errStrs = append(errStrs, e.Error())
	}
	sort.Strings(errStrs)

	// 0
	tchek.AreEqual(t, "multiple error messages", len(errStrs), len(tests))

	if len(errStrs) == len(tests) {
		for i, test := range tests {
			tchek.AreEqual(t, test.name, test.msg, errStrs[i])
		}
	}

}

func TestInstanceDistributor(t *testing.T) {

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

	// app.Registry = jsonapi.NewMockRegistry()

	// app.Get("/mocktypes3", KernelGetCollection, GatePublic)

	return app
}
