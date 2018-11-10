package karigo

import (
	"github.com/kkaribu/jsonapi"
)

// Access ...
type Access struct{}

// Ready ...
func (a *Access) Ready() {}

// WillRead ...
func (a *Access) WillRead(lock string) {}

// WillWrite ...
func (a *Access) WillWrite(lock string) {}

// Release ...
func (a *Access) Release(lock string) {}

// ReleasePartially ...
func (a *Access) ReleasePartially(lock string, keep []string) {}

// ReleaseAll ...
func (a *Access) ReleaseAll() {}

// End ...
func (a *Access) End() NTx {
	return NTx{}
}

// GetString ...
func (a *Access) GetString(key string) string {
	return ""
}

// GetInt ...
func (a *Access) GetInt(key string) int {
	return 0
}

// GetStrings ...
func (a *Access) GetStrings(key string) []string {
	return []string{}
}

// GetInts ...
func (a *Access) GetInts(key string) []int {
	return []int{}
}

// GetManyStructs ...
func (a *Access) GetManyStructs(key string, fields []string, filter *jsonapi.Condition, sort []string, pageSize, pageNumber int, v interface{}) {
}

// SetString ...
func (a *Access) SetString(key string, value string) {}

// SetInt ...
func (a *Access) SetInt(key string, value int) {}

// AddToManyRel ...
func (a *Access) AddToManyRel(key string, id string) {}

// SetToManyRel ...
func (a *Access) SetToManyRel(key string, ids []string) {}

// RemoveToManyRel ...
func (a *Access) RemoveToManyRel(key string, id string) {}
