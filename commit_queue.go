package karigo

import (
	"sync"
)

// commitQueue ...
type commitQueue struct {
	sync.Mutex

	api API
}

// Execute ...
func (c *commitQueue) Execute(f Action) (*Access, error) {
	acc := &Access{}
	err := f.Execute(acc)
	return acc, err
}
