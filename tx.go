package karigo

// Tx ...
type Tx interface {
	Commit() error
	Rollback() error
}

// NTx ...
type NTx struct {
	id     uint
	values []map[string]interface{}
	ops    map[string]interface{}
}

// Set attribute
// Set to-one rel
// Add to-many rel
// Remove to-many rel
// Set to-many rel
