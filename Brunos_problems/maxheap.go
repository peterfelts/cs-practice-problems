package bruno

type MaxHeap []int

// implement the heap interface
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	tmp := *h
	n := len(tmp)
	x := tmp[n-1] // get the last item in the slice
	*h = tmp[0 : n-1]
	return x
}

// implement the sort interface
func (h MaxHeap) Less(a, b int) bool {
	return h[a] > h[b] // max heap
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}
