package bruno

import "sync"

// 7. Implement Queue using Stacks (Data Structures)
// Problem: Implement a queue using two stacks. You need to implement the standard queue
// operations: enqueue (push), dequeue (pop), peek, and empty.

type Queue struct {
	mu     sync.RWMutex // let's make this thread safe, just for un
	stackA []int        // the first stack that receives incoming items
	stackB []int        // the reverse-order statck that outgoing items are popped from
}

func InitQueue() *Queue {
	return &Queue{
		stackA: []int{},
		stackB: []int{},
	}
}

func (q *Queue) Enqueue(item int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.stackA = append(q.stackA, item)

}
func (q *Queue) Dequeue() int {

	q.mu.Lock()
	defer q.mu.Unlock()

	q.transfer()

	// pop the last item from stack B
	item := q.stackB[len(q.stackB)-1]
	q.stackB = q.stackB[:len(q.stackB)-1]
	return item
}

func (q *Queue) Peek() (int, bool) {

	q.mu.Lock()
	defer q.mu.Unlock()

	q.transfer()
	if len(q.stackB) == 0 {
		return 0, false
	}
	return q.stackB[len(q.stackB)-1], true
}

func (q *Queue) Empty() bool {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return len(q.stackA) == 0 && len(q.stackB) == 0
}

func (q *Queue) transfer() {

	// if stack B is empty, copy items from stack A until it is empty
	if len(q.stackB) == 0 {
		for {
			if len(q.stackA) == 0 {
				break
			}
			item := q.stackA[len(q.stackA)-1]     // get the last item in the first stack
			q.stackA = q.stackA[:len(q.stackA)-1] // remove the last item in the frist stack
			q.stackB = append(q.stackB, item)     // put the item into the second stack
		}
	}
}
