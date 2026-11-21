package two_pointers

import "testing"

func TestMiddleOfLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{"documented odd-length example", []int{0, 1, 2, 3, 4}, 2},
		{"documented even-length example", []int{0, 1, 2, 3, 4, 5}, 3},
		{"single node", []int{7}, 7},
		{"single zero", []int{0}, 0},
		{"two nodes return second", []int{8, 3}, 3},
		{"three nodes", []int{9, 4, 8}, 4},
		{"four nodes return second middle", []int{10, 20, 30, 40}, 30},
		{"negative values", []int{-5, -4, -3, -2, -1}, -3},
		{"zero in the middle", []int{8, -2, 0, 9, 4}, 0},
		{"repeated values with distinct middle", []int{5, 5, 9, 5, 5}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *Node
			for i := len(tt.values) - 1; i >= 0; i-- {
				head = &Node{val: tt.values[i], next: head}
			}

			if got := middleOfLinkedList(head); got != tt.expected {
				t.Errorf("middleOfLinkedList(%v) = %d, want %d", tt.values, got, tt.expected)
			}
		})
	}
}
