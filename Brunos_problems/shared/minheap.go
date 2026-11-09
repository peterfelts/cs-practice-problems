package shared

type MinHeap[T any] struct {
	Items   []T
	Compare func(a, b T) bool
}

// implement the heap interface
func (h *MinHeap[T]) Push(x any) {
	h.Items = append(h.Items, x.(T))
}

func (h *MinHeap[T]) Pop() any {
	tmp := *h
	n := len(tmp.Items)
	x := tmp.Items[n-1] // get the last item in the slice
	h.Items = tmp.Items[0 : n-1]
	return x
}

// implement the sort interface
func (h MinHeap[T]) Less(a, b int) bool {
	return h.Compare(h.Items[a], h.Items[b]) // min heap
}

func (h MinHeap[T]) Len() int {
	return len(h.Items)
}

func (h MinHeap[T]) Swap(a, b int) {
	h.Items[a], h.Items[b] = h.Items[b], h.Items[a]
}
