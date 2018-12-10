package karigo

import "time"

// Log ...
type Log struct {
	sources        []Source
	seq            uint
	log            []Entry
	activeAccesses []*Access
	// ongoingTx      []FTx
	// tempLog []
}

// // NewLog ...
// func NewLog() *Log {
// 	log := &Log{
// 		seq: 0,
// 		log: []Entry{},
// 	}

// 	return log
// }

// Entry ...
type Entry struct {
	seq uint
	ops []Op
}

// Op ...
// Set attribute
// Set to-one rel
// Add to-many rel
// Remove to-many rel
// Set to-many rel
type Op struct {
	set   string
	field string
	op    string // set, add, rem
	val   string
}

// Run ...
func (l *Log) Run() error {
	alive := true

	for alive {
		time.Sleep(2 * time.Second)
	}

	return nil
}

// NewAccess ...
func (l *Log) NewAccess() *Access {
	return &Access{
		ops: []Op{},
	}
}

// Execute ...
func (l *Log) Execute(f FTx) error {
	acc := l.NewAccess()
	l.activeAccesses = append(l.activeAccesses, acc)
	err := f(acc)
	acc.done = true
	return err
}

// // ReadAsync ...
// func (l *Log) ReadAsync(f func(*Access) error) error {
// 	acc := l.NewAccess()
// 	err := f(acc)
// 	if err != nil {
// 		return err
// 	}

// 	l.log = append(l.log, Entry{
// 		seq: l.seq,
// 		ops: acc.ops,
// 	})

// 	return err
// }

// // LastSequence ...
// func (l *Log) LastSequence() uint { return l.seq }
