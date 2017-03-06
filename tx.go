package karigo

// Tx ...
type Tx interface {
	Commit() error
	Rollback() error
}
