package karigo

import (
	"sync"
	"time"
)

// commitQueue ...
type commitQueue struct {
	sync.Mutex
	inbox chan Action

	sources []struct {
		src      Source
		versions map[string]map[string]uint
	}

	// Ongoing transactions
	activeAccesses []*Access
}

// Op ...
// Set attribute
// Set to-one rel
// Add to-many rel
// Remove to-many rel
// Set to-many rel
type Op struct {
	set   string
	field string
	op    string // set, add, rem
	val   string
}

// run ...
func (c *commitQueue) run() {
	for {
		time.Sleep(10 * time.Second)

		c.Lock()

		// for a := range c.activeAccesses {

		// }

		c.Unlock()
	}
}

func sourceWorker(s Source, c *commitQueue) {
	var (
	// curr uint
	// acc  *Access
	)

	for {
		c.Lock()

		c.Unlock()
	}
}

// newAccess ...
func (c *commitQueue) newAccess() *Access {
	acc := &Access{
		ops: []Op{},
	}

	c.Lock()
	c.activeAccesses = append(c.activeAccesses, acc)
	c.Unlock()

	return acc
}

// Execute ...
func (c *commitQueue) Execute(f Action) (*Access, error) {
	acc := c.newAccess()
	err := f.Execute(acc)
	acc.done = true
	return acc, err
}
