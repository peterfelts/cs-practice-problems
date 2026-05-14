package bruno

type MinHeap []int

// implement the heap interface
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	tmp := *h
	n := len(tmp)
	x := tmp[n-1] // get the last item in the slice
	*h = tmp[0 : n-1]
	return x
}

// implement the sort interface
func (h MinHeap) Less(a, b int) bool {
	return h[a] < h[b] // min heap
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}
