package karigo

// Driver ...
type Driver interface{}

// ResProvider ...
type ResProvider interface {
	Get(Key) interface{}
}

// ColProvider ...
type ColProvider interface {
	Get(Query) interface{}
}
