package karigo

import (
	"sync"
)

// NewQueue ...
func NewQueue() *Queue {
	queue := &Queue{}

	queue.actions = make([]Action, 0, 100)

	return queue
}

// Queue ...
type Queue struct {
	actions  []Action
	position uint8
	length   uint8
	lock     sync.Mutex

	lastTx []*Tx
}

// Run ...
func (q *Queue) Run() {
	for {
		q.lock.Lock()

		if q.length > 0 {
			// q.actions[q.position].Execute(ctx, sw)
		}

		q.lock.Unlock()
	}
}

// Execute ...
func (q *Queue) Execute(a Action) {
	q.lock.Lock()

	pos := q.position + q.length

	q.actions[pos] = a

	q.lock.Unlock()

}

// func (q *Queue) Lock(lock string) uint {
// 	if l, ok := q.locks[lock]; ok {
// 		return l
// 	}
//
// 	return 0
// }

// func (q *Queue) Insert(tx []Tx) error {
// 	return nil
// }
