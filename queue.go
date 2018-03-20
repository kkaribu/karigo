package karigo

import (
	"sync"
)

func NewQueue() *Queue {
	queue := &Queue{}

	queue.actions = make([]Action, 0, 100)

	return queue
}

type Queue struct {
	actions  []Action
	position int8
	length   uint8
	lock     sync.Mutex

	locks  map[string]uint
	lastTx []*Tx
}

func (q *Queue) Run() {
	for {
		q.lock.Lock()

		if q.actions.length > 0 {
			q.actions[position].Execute(ctx, sw)
		}

		q.lock.Unlock()
	}
}

func (q *Queue) Execute(a Action) {
	q.lock.Lock()

	pos := q.position + q.length

	queue.actions[pos]

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
