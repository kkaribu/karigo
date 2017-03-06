package karigo

import "github.com/kkaribu/jsonapi"

// A Gate decides whether a request is allowed to continue or not.
type Gate func(*Ctx) bool

// checkGates...
func (a *App) checkGates(ctx *Ctx) {
	for _, gate := range a.Gates[ctx.Method+" "+ctx.URL.Route] {
		if !gate(ctx) {
			panic(jsonapi.NewErrForbidden())
		}
	}
}

/*
 * DEFAULT GATES
 */

// GatePublic ...
func GatePublic(*Ctx) bool {
	return true
}

// GateLoggedIn ...
func GateLoggedIn(ctx *Ctx) bool {
	return ctx.ID != ""
}
