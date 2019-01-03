package karigo

import (
	"fmt"
)

// DefaultSource ...
type DefaultSource struct {
}

// Field ...
func (d *DefaultSource) Field(qry Query) (string, error) {
	if qry.ID == "" {
		return "", fmt.Errorf("karigo: a query supplied to Field needs a non empty ID")
	}

	return "", nil
}

// Fields ...
func (d *DefaultSource) Fields(qry Query) ([][]string, error) {
	if qry.ID == "" {
		return nil, fmt.Errorf("karigo: a query supplied to Field needs a non empty ID")
	}

	return nil, nil
}

// Apply ...
func (d *DefaultSource) Apply(ops []Op) error {
	return nil
}
