package karigo

import "github.com/kkaribu/jsonapi"

// Store ...
type Store interface {
	// Connection management
	Close()

	// Transaction
	Begin() (Tx, error)

	// Resource manipulation
	SelectCollection(tx Tx, resType string, params *jsonapi.Params, c jsonapi.Collection) error
	SelectResource(tx Tx, resType, resID string, params *jsonapi.Params, r jsonapi.Resource) error
	SelectInclusions(tx Tx, originType, originID string, params *jsonapi.Params) ([]jsonapi.Resource, error)
	InsertResource(tx Tx, r jsonapi.Resource) error
	UpdateResource(tx Tx, resType, resID string, updates map[string]interface{}) error
	DeleteResource(tx Tx, resType, resID string) error
	ResourceExists(tx Tx, resType, resID string) (bool, error)

	// Relationship manipulation
	SelectRelationship(tx Tx, resType string, resID string, rel jsonapi.Rel) (string, error)
	SelectRelationships(tx Tx, resType string, resID string, rel jsonapi.Rel) ([]string, error)
	UpdateRelationship(tx Tx, resType string, resID string, rel jsonapi.Rel, relID string) error
	UpdateRelationships(tx Tx, resType string, resID string, rel jsonapi.Rel, relIDs []string) error
	InsertRelationships(tx Tx, resType string, resID string, rel jsonapi.Rel, relIDs []string) error
	DeleteRelationship(tx Tx, resType, resID string, rel jsonapi.Rel, relID string) error
	DeleteRelationships(tx Tx, resType, resID string, rel jsonapi.Rel, relIDs []string) error
	DeleteAllRelationships(tx Tx, resType, resID string, rel jsonapi.Rel) error

	// Database management
	SetRegistry(reg *jsonapi.Registry)
	Open(user string, pw string, host string, dbName string) error
	SelectResourceTables(tx Tx) ([]string, error)
	SelectRelationshipTables(tx Tx) ([]string, error)
	SelectColumns(tx Tx, resType string) ([]map[string]string, error)
	DrainDatabase(tx Tx) error
	SyncDatabase(tx Tx, reg *jsonapi.Registry, verbose bool) error
}
