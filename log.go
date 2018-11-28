package karigo

// Log ...
type Log struct {
	drivers []Driver
	seq     uint
	log     []Entry
}

// NewLog ...
func NewLog() *Log {
	log := &Log{
		seq: 0,
		log: []Entry{},
	}

	return log
}

// Entry ...
type Entry struct {
	seq    uint
	id     string
	values []map[string]interface{}
	ops    []Op
}

// Op ...
// Set attribute
// Set to-one rel
// Add to-many rel
// Remove to-many rel
// Set to-many rel
type Op struct {
	field string
	op    string // set, add, rem
	val   interface{}
}

// NewAccess ...
func (l *Log) NewAccess() *Access {
	return &Access{
		drivers: []Driver{},
		ops:     []Op{},
	}
}

// Execute ...
func (l *Log) Execute(f func(*Access) error) error {
	acc := l.NewAccess()
	err := f(acc)
	return err
}

// ReadAsync ...
func (l *Log) ReadAsync(f func(*Access) error) error {
	acc := l.NewAccess()
	err := f(acc)
	if err != nil {
		return err
	}

	l.log = append(l.log, Entry{
		seq: l.seq,
		id:  "", // TODO
		ops: acc.ops,
	})

	return err
}

// LastSequence ...
func (l *Log) LastSequence() uint { return l.seq }
