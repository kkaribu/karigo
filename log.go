package karigo

// Log ...
type Log struct {
	seq int
	log []Entry
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
	return &Access{}
}

// Execute ...
func (l *Log) Execute(func(*Access)) error { return nil }

// ReadAsync ...
func (l *Log) ReadAsync(func(*Access)) error { return nil }

// LastSequence ...
func (l *Log) LastSequence() int { return l.seq }
