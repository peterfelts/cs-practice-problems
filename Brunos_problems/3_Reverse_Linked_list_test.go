package bruno

import (
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"multiple elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"single element", []int{42}, []int{42}},
		{"already reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := listFromSlice(tt.input)
			result := reverseLinkedList(input)
			if got := sliceFromList(result); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reverseLinkedList(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
