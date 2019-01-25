package karigo

import (
	"github.com/kkaribu/jsonapi"
)

// API ...
type API struct {
	Types       []jsonapi.Type
	CreateFuncs map[string]Action
	UpdateFuncs map[string]Action
	DeleteFuncs map[string]Action
}
