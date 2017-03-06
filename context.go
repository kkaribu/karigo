package karigo

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/kkaribu/jsonapi"
	"gopkg.in/dgrijalva/jwt-go.v2"
)

// NewCtx creates a Ctx and returns it.
func NewCtx(a *App, w ResponseWriter, r *http.Request) *Ctx {
	return &Ctx{
		App: a,

		W: w,

		URL: &jsonapi.URL{
			Params: &jsonapi.Params{},
		},
	}
}

// A Ctx is passed around during a request for storing temporary data
// related to the request.
type Ctx struct {
	sync.Mutex `json:"-"`

	App   *App  `json:"-"`
	Store Store `json:"-"`
	Tx    Tx

	// Log
	Log []string

	// Request
	W      ResponseWriter `json:"-"`
	Method string
	URL    *jsonapi.URL
	Body   []byte

	// User
	JWT    *jwt.Token `json:"-"`
	ID     string
	Groups []string

	// Document
	Doc *jsonapi.Document
}

// AddToLog adds a record to the context's log.
func (ctx *Ctx) AddToLog(l string) {
	ctx.Lock()
	defer ctx.Unlock()

	if l != "" {
		ctx.Log = append(ctx.Log, fmt.Sprintf(l))
	}
}

// SaveLog ...
func (ctx *Ctx) SaveLog() {
	var data string

	for _, l := range ctx.Log {
		data = l + "\n"
	}
	data = data[:len(data)-1] // Remove last \n

	fmt.Printf("%s\n\n", data)

}
