package karigo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kkaribu/jsonapi"
)

// Switch ...
type Switch struct {
	engine *Queue
	// locks  map[string]Lock
}

func (t *Switch) Reserve(key string) error {
	return nil
}

func (t *Switch) ReserveRO(key string) error {
	return nil
}

// You cannot reserver anything else. You may now read and update.
func (t *Switch) Ready() {
}

func (t *Switch) NarrowDown(let string, keep []string) error {
	for _, k := range keep {
		if !strings.HasPrefix(k, let) {
			return errors.New(fmt.Sprintf("engine: %s isn't an inner lock of %s", k, let))
		}
	}

	return nil
}

func (t *Switch) GetCol(typ string, fields []string, pagination [2]int, filters map[string][]string, sort []string) error {
	return nil
}

func (t *Switch) GetRes(typ, id string, fields []string) error {
	return nil
}

func (t *Switch) Insert(res jsonapi.Resource) error {
	return nil
}

func (t *Switch) Update() error {
	return nil
}

func (t *Switch) Delete() error {
	return nil
}
