package karigo

// Tx ...
type Tx interface {
	Commit() error
	Rollback() error
}

// FTx ...
type FTx func(*Access) error
