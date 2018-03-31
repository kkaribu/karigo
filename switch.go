package karigo

import (
	"fmt"
	"strings"

	"github.com/kkaribu/jsonapi"
)

// Switch ...
type Switch struct {
	engine *Queue
	locksR map[string]Lock
	locksW map[string]Lock
}

// ReserveR ...
func (t *Switch) ReserveR(key string) error {
	return nil
}

// ReserveW ...
func (t *Switch) ReserveW(key string) error {
	return nil
}

// Ready ...
// You cannot reserve anything else. You may now read and update.
func (t *Switch) Ready() {
}

// NarrowDown ...
func (t *Switch) NarrowDown(let string, keep []string) error {
	for _, k := range keep {
		if !strings.HasPrefix(k, let) {
			return fmt.Errorf("karigo: %s isn't an inner lock of %s", k, let)
		}
	}

	return nil
}

// GetCol ...
func (t *Switch) GetCol(typ string, fields []string, pagination [2]int, filters map[string][]string, sort []string) (jsonapi.Collection, error) {
	// var col jsonapi.Collection
	//
	// if nf != nil {
	// 	t.NarrowDown("typ", nf(col))
	// } else {
	// 	t.Clear(typ)
	// }

	return nil, nil
}

// GetRes ...sort
func (t *Switch) GetRes(typ, id string, fields []string) (jsonapi.Resource, error) {
	return nil, nil
}

// Insert ...
func (t *Switch) Insert(res jsonapi.Resource) error {
	return nil
}

// Update ...
func (t *Switch) Update() error {
	return nil
}

// Delete ...
func (t *Switch) Delete() error {
	return nil
}
