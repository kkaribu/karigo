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

// Execute ...
func (l *Log) Execute(func(*Access)) error { return nil }

// Read ...
func (l *Log) Read(func(*Access)) error { return nil }

// MethodName ...
func (l *Log) MethodName() {}

// LastSequence ...
func (l *Log) LastSequence() int { return l.seq }

// MethodName3 ...
func (l *Log) MethodName3() {}

// MethodName4 ...
func (l *Log) MethodName4() {}
