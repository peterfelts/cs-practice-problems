package two_pointers

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		n        int
		expected []int
	}{
		{"documented tail removal", []int{1, 2, 3, 4}, 1, []int{1, 2, 3}},
		{"documented interior removal", []int{1, 2, 3, 4}, 2, []int{1, 2, 4}},
		{"documented head removal", []int{1, 2, 3, 4}, 4, []int{2, 3, 4}},
		{"single node", []int{7}, 1, nil},
		{"two nodes remove head", []int{8, 3}, 2, []int{3}},
		{"two nodes remove tail", []int{8, 3}, 1, []int{8}},
		{"odd-length middle removal", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
		{"remove second node", []int{1, 2, 3, 4, 5}, 4, []int{1, 3, 4, 5}},
		{"repeated values", []int{5, 2, 5, 3}, 2, []int{5, 2, 3}},
		{"negative values and zero", []int{-3, 0, -1, 2}, 3, []int{-3, -1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *Node
			for i := len(tt.values) - 1; i >= 0; i-- {
				head = &Node{val: tt.values[i], next: head}
			}

			got := removeNthFromEnd(head, tt.n)
			for i, want := range tt.expected {
				if got == nil {
					t.Fatalf("result ended at index %d, want %v", i, tt.expected)
				}
				if got.val != want {
					t.Errorf("result at index %d = %d, want %d", i, got.val, want)
				}
				got = got.next
			}
			if got != nil {
				t.Errorf("result has extra nodes after %d expected nodes", len(tt.expected))
			}
		})
	}
}
