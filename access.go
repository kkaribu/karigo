package karigo

import (
	"github.com/kkaribu/jsonapi"
)

// Access ...
type Access struct {
	deps     []Query
	affected []jsonapi.Resource
	ops      []Op
	done     bool
}

// LockR ...
func (a *Access) LockR() {}

// LockW ...
func (a *Access) LockW() {}

// Ready ...
func (a *Access) Ready() {}

// Do ...
func (a *Access) Do(op Op) {}

// Release ...
func (a *Access) Release(query Query, keep ...[]string) {}

// GetVal ...
func (a *Access) GetVal(query Query) interface{} { return nil }

// GetVals ...
func (a *Access) GetVals(query Query) []interface{} { return []interface{}{} }

// GetRes ...
func (a *Access) GetRes(query Query) map[string]interface{} { return nil }

// GetCol ...
func (a *Access) GetCol(query Query) []map[string]interface{} { return nil }

// GetInclusions ...
func (a *Access) GetInclusions(query Query, rels, fields []string) map[string]map[string]interface{} {
	return map[string]map[string]interface{}{}
}

// Count ...
func (a *Access) Count(query Query) int {
	return 0
}
