package bruno

import "testing"

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := InitQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if got := q.Dequeue(); got != 1 {
		t.Errorf("Dequeue() = %d, want 1", got)
	}
	if got := q.Dequeue(); got != 2 {
		t.Errorf("Dequeue() = %d, want 2", got)
	}
	if got := q.Dequeue(); got != 3 {
		t.Errorf("Dequeue() = %d, want 3", got)
	}
}

func TestQueue_FIFOOrder(t *testing.T) {
	q := InitQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	if got := q.Dequeue(); got != 10 {
		t.Errorf("Dequeue() = %d, want 10", got)
	}
	q.Enqueue(30)
	if got := q.Dequeue(); got != 20 {
		t.Errorf("Dequeue() = %d, want 20", got)
	}
	if got := q.Dequeue(); got != 30 {
		t.Errorf("Dequeue() = %d, want 30", got)
	}
}

func TestQueue_Peek(t *testing.T) {
	q := InitQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	got, ok := q.Peek()
	if !ok {
		t.Fatal("Peek() returned ok=false on non-empty queue, want true")
	}
	if got != 1 {
		t.Errorf("Peek() = %d, want 1", got)
	}
	// Peek should not remove the item
	got, _ = q.Peek()
	if got != 1 {
		t.Errorf("Peek() after second call = %d, want 1 (should not remove item)", got)
	}
	// Dequeue should still return the peeked item
	if got := q.Dequeue(); got != 1 {
		t.Errorf("Dequeue() after Peek() = %d, want 1", got)
	}
	// Peek should now reflect the new front
	got, ok = q.Peek()
	if !ok {
		t.Fatal("Peek() returned ok=false after one dequeue, want true")
	}
	if got != 2 {
		t.Errorf("Peek() after Dequeue() = %d, want 2", got)
	}
}

func TestQueue_Peek_Empty(t *testing.T) {
	q := InitQueue()
	_, ok := q.Peek()
	if ok {
		t.Error("Peek() on empty queue returned ok=true, want false")
	}
}

func TestQueue_Empty(t *testing.T) {
	q := InitQueue()
	if !q.Empty() {
		t.Error("Empty() = false on a new queue, want true")
	}
	q.Enqueue(1)
	if q.Empty() {
		t.Error("Empty() = true after Enqueue, want false")
	}
	q.Dequeue()
	if !q.Empty() {
		t.Error("Empty() = false after dequeuing all items, want true")
	}
}
