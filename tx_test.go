package karigo

type MockTx struct {
	CommitFunc   func() error
	RollbackFunc func() error
}

func (m MockTx) Commit() error {
	if m.CommitFunc == nil {
		panic("function Commit in MockTx not implemented")
	}
	return m.CommitFunc()
}

func (m MockTx) Rollback() error {
	if m.RollbackFunc == nil {
		panic("function Rollback in MockTx not implemented")
	}
	return m.RollbackFunc()
}
