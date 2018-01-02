package karigo

import "github.com/kkaribu/jsonapi"

// Store ...
type Store interface {
	// Connection management
	Open(driver, url string) error
	Close()

	// Transaction
	Begin() (Tx, error)

	// Resource manipulation
	SelectCollection(tx Tx, resType string, from jsonapi.FromFilter, params *jsonapi.Params, c jsonapi.Collection) error
	SelectResource(tx Tx, resType, resID string, from jsonapi.FromFilter, params *jsonapi.Params, r jsonapi.Resource) error
	SelectInclusions(tx Tx, originType, originID string, from jsonapi.FromFilter, params *jsonapi.Params) ([]jsonapi.Resource, error)
	InsertResource(tx Tx, r jsonapi.Resource) error
	UpdateResource(tx Tx, resType, resID string, updates map[string]interface{}) error
	DeleteResource(tx Tx, resType, resID string) error
	ResourceExists(tx Tx, resType, resID string) (bool, error)

	// Relationship manipulation
	SelectRelationship(tx Tx, resType, resID, relName string) (string, error)
	SelectRelationships(tx Tx, resType, resID, relName string) ([]string, error)
	UpdateRelationship(tx Tx, resType, resID, relName, relID string) error
	UpdateRelationships(tx Tx, resType, resID, relName string, relIDs []string) error
	InsertRelationships(tx Tx, resType, resID, relName string, relIDs []string) error
	DeleteRelationship(tx Tx, resType, resID, relID, relName string) error
	DeleteRelationships(tx Tx, resType, resID, relName string, relIDs []string) error
	DeleteAllRelationships(tx Tx, resType, resID, relName string) error

	// Other
	CountCollectionSize(tx Tx, resType string, from jsonapi.FromFilter, params *jsonapi.Params) (int, error)

	// Database management
	SetRegistry(reg *jsonapi.Registry)
	SelectResourceTables(tx Tx) ([]string, error)
	SelectRelationshipTables(tx Tx) ([]string, error)
	SelectColumns(tx Tx, resType string) ([]map[string]string, error)
	CreateResourceTable(tx Tx, typ jsonapi.Type) error
	CreateRelationshipTable(tx Tx, rel jsonapi.Rel) error
	AddColumn(tx Tx, resType string, attr jsonapi.Attr) error
	DropTable(tx Tx, resType string) error
	DropColumn(tx Tx, resType, colName string) error
	DrainDatabase(tx Tx) error
	SyncDatabase(tx Tx, reg *jsonapi.Registry, verbose, apply bool) error
}
