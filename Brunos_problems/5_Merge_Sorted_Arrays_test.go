package bruno

import (
	"reflect"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		m, n     int
		expected []int
	}{
		{
			name:     "example case",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			nums2:    []int{2, 5, 6},
			m:        3, n: 3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "nums2 all larger",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			nums2:    []int{4, 5, 6},
			m:        3, n: 3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "nums2 all smaller",
			nums1:    []int{4, 5, 6, 0, 0, 0},
			nums2:    []int{1, 2, 3},
			m:        3, n: 3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "nums1 empty",
			nums1:    []int{0, 0, 0},
			nums2:    []int{1, 2, 3},
			m:        0, n: 3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "nums2 empty",
			nums1:    []int{1, 2, 3},
			nums2:    []int{},
			m:        3, n: 0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "single element each",
			nums1:    []int{1, 0},
			nums2:    []int{2},
			m:        1, n: 1,
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeSortedArrays(tt.nums1, tt.nums2, tt.m, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mergeSortedArrays(%v, %v, %d, %d) = %v, want %v",
					tt.nums1, tt.nums2, tt.m, tt.n, result, tt.expected)
			}
		})
	}
}
