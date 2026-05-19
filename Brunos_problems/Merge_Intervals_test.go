package bruno

import (
	"reflect"
	"testing"
)

func Test_MergeIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals []interval
		expected  []interval
	}{
		{
			name: "basic overlapping unsorted",
			intervals: []interval{
				{13, 17},
				{0, 1},
				{7, 9},
				{3, 12},
				{4, 9},
			},
			expected: []interval{
				{0, 1},
				{3, 12},
				{13, 17},
			},
		},
		{
			name: "no overlaps already sorted",
			intervals: []interval{
				{1, 2},
				{5, 7},
				{10, 15},
			},
			expected: []interval{
				{1, 2},
				{5, 7},
				{10, 15},
			},
		},
		{
			name: "all intervals overlap into one",
			intervals: []interval{
				{1, 3},
				{2, 5},
				{2, 100},
			},
			expected: []interval{
				{1, 100},
			},
		},
		{
			name: "single interval",
			intervals: []interval{
				{5, 10},
			},
			expected: []interval{
				{5, 10},
			},
		},
		{
			name: "two non-overlapping intervals",
			intervals: []interval{
				{1, 3},
				{5, 7},
			},
			expected: []interval{
				{1, 3},
				{5, 7},
			},
		},
		{
			name: "two overlapping intervals",
			intervals: []interval{
				{1, 5},
				{3, 8},
			},
			expected: []interval{
				{1, 8},
			},
		},
		{
			name: "touching intervals (share endpoint)",
			intervals: []interval{
				{1, 5},
				{5, 10},
			},
			expected: []interval{
				{1, 10},
			},
		},
		{
			name: "one interval completely contained in another",
			intervals: []interval{
				{1, 10},
				{3, 7},
			},
			expected: []interval{
				{1, 10},
			},
		},
		{
			name: "identical intervals",
			intervals: []interval{
				{4, 8},
				{4, 8},
				{4, 8},
			},
			expected: []interval{
				{4, 8},
			},
		},
		{
			name: "point intervals (left == right)",
			intervals: []interval{
				{3, 3},
				{5, 5},
				{3, 3},
			},
			expected: []interval{
				{3, 3},
				{5, 5},
			},
		},
		{
			name: "multiple disjoint groups",
			intervals: []interval{
				{1, 4},
				{2, 6},
				{10, 12},
				{11, 15},
				{20, 25},
			},
			expected: []interval{
				{1, 6},
				{10, 15},
				{20, 25},
			},
		},
		{
			name: "unsorted input merges into two",
			intervals: []interval{
				{15, 20},
				{1, 5},
				{3, 7},
				{16, 18},
			},
			expected: []interval{
				{1, 7},
				{15, 20},
			},
		},
		{
			name: "large gap between intervals",
			intervals: []interval{
				{0, 1},
				{1000, 2000},
			},
			expected: []interval{
				{0, 1},
				{1000, 2000},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := mergeIntervals(tt.intervals)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("mergeIntervals() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
