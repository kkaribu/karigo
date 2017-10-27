package karigo

type MockTx struct {
	CommitFunc   func() error
	RollbackFunc func() error
}

func (m MockTx) Commit() error {
	if m.CommitFunc == nil {
		return nil
	}
	return m.CommitFunc()
}

func (m MockTx) Rollback() error {
	if m.RollbackFunc == nil {
		return nil
	}
	return m.RollbackFunc()
}
