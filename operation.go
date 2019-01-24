package karigo

// Op ...
// Set attribute
// Set to-one rel
// Add to-many rel
// Remove to-many rel
// Set to-many rel
type Op struct {
	Set   string
	ID    string
	Field string
	Op    string // set, add, rem
	Val   interface{}
}
