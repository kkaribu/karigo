package karigo

// Source ...
type Source interface {
	Field(Query) (string, error)
	Fields(Query) ([][]string, error)
	Apply([]Op) error
}
