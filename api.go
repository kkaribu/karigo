package karigo

import (
	"github.com/kkaribu/jsonapi"
)

// API ...
type API struct {
	Types         []jsonapi.Type
	GateFuncs     map[string]Action
	ValidateFuncs map[string]Action
	CreateFuncs   map[string]Action
	UpdateFuncs   map[string]Action
	DeleteFuncs   map[string]Action
}
