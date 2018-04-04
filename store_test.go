package karigo

import "github.com/kkaribu/jsonapi"

// MockStore ...
type MockStore struct {
	OpenFunc  func(driver, host, db, user, pw string, opts map[string]string) error
	URLFunc   func() string
	CloseFunc func()

	BeginFunc func() (Tx, error)

	SelectCollectionFunc func(tx Tx, resType string, from jsonapi.FromFilter, params *jsonapi.Params, c jsonapi.Collection) error
	SelectResourceFunc   func(tx Tx, resType, resID string, from jsonapi.FromFilter, params *jsonapi.Params, r jsonapi.Resource) error
	SelectInclusionsFunc func(tx Tx, originType, originID string, from jsonapi.FromFilter, params *jsonapi.Params) ([]jsonapi.Resource, error)
	InsertResourceFunc   func(tx Tx, r jsonapi.Resource) error
	UpdateResourceFunc   func(tx Tx, resType, resID string, updates map[string]interface{}) error
	DeleteResourceFunc   func(tx Tx, resType, resID string) error
	ResourceExistsFunc   func(tx Tx, resType, resID string) (bool, error)

	SelectRelationshipFunc     func(tx Tx, resType string, resID, relName string) (string, error)
	SelectRelationshipsFunc    func(tx Tx, resType string, resID, relName string) ([]string, error)
	UpdateRelationshipFunc     func(tx Tx, resType string, resID, relName string, relID string) error
	UpdateRelationshipsFunc    func(tx Tx, resType string, resID, relName string, relIDs []string) error
	InsertRelationshipsFunc    func(tx Tx, resType string, resID, relName string, relIDs []string) error
	DeleteRelationshipFunc     func(tx Tx, resType, resID, relID, relName string) error
	DeleteRelationshipsFunc    func(tx Tx, resType, resID, relName string, relIDs []string) error
	DeleteAllRelationshipsFunc func(tx Tx, resType, resID, relName string) error

	CountCollectionSizeFunc func(tx Tx, resType string, from jsonapi.FromFilter, params *jsonapi.Params) (int, error)

	SetRegistryFunc              func(reg *jsonapi.Registry)
	SelectResourceTablesFunc     func(tx Tx) ([]string, error)
	SelectRelationshipTablesFunc func(tx Tx) ([]string, error)
	SelectColumnsFunc            func(tx Tx, resType string) ([]map[string]string, error)
	CreateResourceTableFunc      func(tx Tx, typ jsonapi.Type) error
	CreateRelationshipTableFunc  func(tx Tx, rel jsonapi.Rel) error
	AddColumnFunc                func(tx Tx, resType string, attr jsonapi.Attr) error
	DropTableFunc                func(tx Tx, resType string) error
	DropColumnFunc               func(tx Tx, resType, colName string) error
	DrainDatabaseFunc            func(tx Tx) error
	SyncDatabaseFunc             func(tx Tx, reg *jsonapi.Registry, verbose, apply bool) error

	*jsonapi.Registry
}

func (m MockStore) Open(driver, host, db, user, pw string, opts map[string]string) error {
	if m.OpenFunc == nil {
		panic("function Open in MockStore not implemented")
	}
	return m.OpenFunc(driver, host, db, user, pw, opts)
}

func (m MockStore) URL() string {
	if m.URLFunc == nil {
		panic("function URL in MockStore not implemented")
	}
	return m.URLFunc()
}

func (m MockStore) Close() {
	if m.CloseFunc != nil {
		m.CloseFunc()
	}
}

func (m MockStore) Begin() (Tx, error) {
	if m.BeginFunc == nil {
		return MockTx{}, nil
	}
	return m.BeginFunc()
}

func (m MockStore) SelectCollection(tx Tx, resType string, from jsonapi.FromFilter, params *jsonapi.Params, c jsonapi.Collection) error {
	if m.SelectCollectionFunc == nil {
		panic("function SelectCollection in MockStore not implemented")
	}
	return m.SelectCollectionFunc(tx, resType, from, params, c)
}

func (m MockStore) SelectResource(tx Tx, resType, resID string, from jsonapi.FromFilter, params *jsonapi.Params, r jsonapi.Resource) error {
	if m.SelectResourceFunc == nil {
		panic("function SelectResource in MockStore not implemented")
	}
	return m.SelectResourceFunc(tx, resType, resID, from, params, r)
}

func (m MockStore) SelectInclusions(tx Tx, originType, originID string, from jsonapi.FromFilter, params *jsonapi.Params) ([]jsonapi.Resource, error) {
	if m.SelectInclusionsFunc == nil {
		return nil, nil
	}
	return m.SelectInclusionsFunc(tx, originType, originID, from, params)
}

func (m MockStore) InsertResource(tx Tx, r jsonapi.Resource) error {
	if m.InsertResourceFunc == nil {
		panic("function InsertResource in MockStore not implemented")
	}
	return m.InsertResourceFunc(tx, r)
}

func (m MockStore) UpdateResource(tx Tx, resType, resID string, updates map[string]interface{}) error {
	if m.UpdateResourceFunc == nil {
		panic("function UpdateResource in MockStore not implemented")
	}
	return m.UpdateResourceFunc(tx, resType, resID, updates)
}

func (m MockStore) DeleteResource(tx Tx, resType, resID string) error {
	if m.DeleteResourceFunc == nil {
		panic("function DeleteResource in MockStore not implemented")
	}
	return m.DeleteResourceFunc(tx, resType, resID)
}

func (m MockStore) ResourceExists(tx Tx, resType, resID string) (bool, error) {
	if m.ResourceExistsFunc == nil {
		panic("function ResourceExists in MockStore not implemented")
	}
	return m.ResourceExistsFunc(tx, resType, resID)
}

func (m MockStore) SelectRelationship(tx Tx, resType, resID, relName string) (string, error) {
	if m.SelectRelationshipFunc == nil {
		panic("function SelectRelationship in MockStore not implemented")
	}
	return m.SelectRelationshipFunc(tx, resType, resID, relName)
}

func (m MockStore) SelectRelationships(tx Tx, resType, resID, relName string) ([]string, error) {
	if m.SelectRelationshipsFunc == nil {
		panic("function SelectRelationships in MockStore not implemented")
	}
	return m.SelectRelationshipsFunc(tx, resType, resID, relName)
}

func (m MockStore) UpdateRelationship(tx Tx, resType, resID, relName, relID string) error {
	if m.UpdateRelationshipFunc == nil {
		panic("function UpdateRelationship in MockStore not implemented")
	}
	return m.UpdateRelationshipFunc(tx, resType, resID, relName, relID)
}

func (m MockStore) UpdateRelationships(tx Tx, resType, resID, relName string, relIDs []string) error {
	if m.UpdateRelationshipsFunc == nil {
		panic("function UpdateRelationships in MockStore not implemented")
	}
	return m.UpdateRelationshipsFunc(tx, resType, resID, relName, relIDs)
}

func (m MockStore) InsertRelationships(tx Tx, resType, resID, relName string, relIDs []string) error {
	if m.InsertRelationshipsFunc == nil {
		panic("function InsertRelationships in MockStore not implemented")
	}
	return m.InsertRelationshipsFunc(tx, resType, resID, relName, relIDs)
}

func (m MockStore) DeleteRelationship(tx Tx, resType, resID, relID, relName string) error {
	if m.DeleteRelationshipFunc == nil {
		panic("function DeleteRelationship in MockStore not implemented")
	}
	return m.DeleteRelationshipFunc(tx, resType, resID, relID, relName)
}

func (m MockStore) DeleteRelationships(tx Tx, resType, resID, relName string, relIDs []string) error {
	if m.DeleteRelationshipsFunc == nil {
		panic("function DeleteRelationships in MockStore not implemented")
	}
	return m.DeleteRelationshipsFunc(tx, resType, resID, relName, relIDs)
}

func (m MockStore) DeleteAllRelationships(tx Tx, resType, resID, relName string) error {
	if m.DeleteAllRelationshipsFunc == nil {
		panic("function DeleteAllRelationships in MockStore not implemented")
	}
	return m.DeleteAllRelationshipsFunc(tx, resType, resID, relName)
}

func (m MockStore) CountCollectionSize(tx Tx, resType string, from jsonapi.FromFilter, params *jsonapi.Params) (int, error) {
	if m.CountCollectionSizeFunc == nil {
		return 0, nil
	}
	return m.CountCollectionSizeFunc(tx, resType, from, params)
}

func (m MockStore) SetRegistry(reg *jsonapi.Registry) {
	if m.SetRegistryFunc == nil {
		m.Registry = reg
	}
	m.SetRegistryFunc(reg)
}

func (m MockStore) SelectResourceTables(tx Tx) ([]string, error) {
	if m.SelectResourceTablesFunc == nil {
		panic("function SelectResourceTables in MockStore not implemented")
	}
	return m.SelectResourceTablesFunc(tx)
}

func (m MockStore) SelectRelationshipTables(tx Tx) ([]string, error) {
	if m.SelectRelationshipTablesFunc == nil {
		panic("function SelectRelationshipTables in MockStore not implemented")
	}
	return m.SelectRelationshipTablesFunc(tx)
}

func (m MockStore) SelectColumns(tx Tx, resType string) ([]map[string]string, error) {
	if m.SelectColumnsFunc == nil {
		panic("function SelectColumns in MockStore not implemented")
	}
	return m.SelectColumnsFunc(tx, resType)
}

func (m MockStore) CreateResourceTable(tx Tx, typ jsonapi.Type) error {
	if m.CreateResourceTableFunc == nil {
		panic("function CreateResourceTable in MockStore not implemented")
	}
	return m.CreateResourceTableFunc(tx, typ)
}

func (m MockStore) CreateRelationshipTable(tx Tx, rel jsonapi.Rel) error {
	if m.CreateRelationshipTableFunc == nil {
		panic("function CreateRelationshipTable in MockStore not implemented")
	}
	return m.CreateRelationshipTableFunc(tx, rel)
}

func (m MockStore) AddColumn(tx Tx, resType string, attr jsonapi.Attr) error {
	if m.AddColumnFunc == nil {
		panic("function AddColumn in MockStore not implemented")
	}
	return m.AddColumnFunc(tx, resType, attr)
}

func (m MockStore) DropTable(tx Tx, resType string) error {
	if m.DropTableFunc == nil {
		panic("function DropTable in MockStore not implemented")
	}
	return m.DropTableFunc(tx, resType)
}

func (m MockStore) DropColumn(tx Tx, resType, colName string) error {
	if m.DropColumnFunc == nil {
		panic("function DropColumn in MockStore not implemented")
	}
	return m.DropColumnFunc(tx, resType, colName)
}

func (m MockStore) DrainDatabase(tx Tx) error {
	if m.DrainDatabaseFunc == nil {
		panic("function DrainDatabase in MockStore not implemented")
	}
	return m.DrainDatabaseFunc(tx)
}

func (m MockStore) SyncDatabase(tx Tx, reg *jsonapi.Registry, verbose, apply bool) error {
	if m.SyncDatabaseFunc == nil {
		panic("function SyncDatabase in MockStore not implemented")
	}
	return m.SyncDatabaseFunc(tx, reg, verbose, apply)
}
