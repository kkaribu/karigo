package karigo

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/kkaribu/jsonapi"
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

	App *App `json:"-"`
	Tx  func(acc *Access)

	// Log
	Log []string

	// Request
	W      ResponseWriter `json:"-"`
	Method string
	URL    *jsonapi.URL
	Query  Query
	Key    Key

	// User
	JWT    *jwt.Token `json:"-"`
	ID     string
	Groups []string

	// Payloads
	In  *jsonapi.Payload
	Out *jsonapi.Document
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
		data += l + "\n"
	}
	data = data[:len(data)-1] // Remove last \n

	fmt.Printf("%s\n\n", data)

}
