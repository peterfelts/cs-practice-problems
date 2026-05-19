package bruno

import (
	"testing"
)

func Test_canFinish(t *testing.T) {
	tests := []struct {
		name         string
		numCourses   int
		prerequisites [][]int
		expected     bool
	}{
		{
			name:         "no prerequisites",
			numCourses:   5,
			prerequisites: [][]int{},
			expected:     true,
		},
		{
			name:         "one prerequisite",
			numCourses:   2,
			prerequisites: [][]int{{1, 0}},
			expected:     true,
		},
		{
			name:         "simple two-course cycle",
			numCourses:   2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expected:     false,
		},
		{
			name:         "linear chain of three no cycle",
			numCourses:   3,
			prerequisites: [][]int{{1, 0}, {2, 1}},
			expected:     true,
		},
		{
			name:         "three-course cycle",
			numCourses:   3,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
			expected:     false,
		},
		{
			name:         "diamond DAG no cycle",
			numCourses:   4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:     true,
		},
		{
			name:         "star graph all depend on course 0",
			numCourses:   5,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 0}, {4, 0}},
			expected:     true,
		},
		{
			name:         "multiple disconnected components all acyclic",
			numCourses:   6,
			prerequisites: [][]int{{1, 0}, {3, 2}, {5, 4}},
			expected:     true,
		},
		{
			name:         "multiple components one has cycle",
			numCourses:   6,
			prerequisites: [][]int{{1, 0}, {3, 2}, {2, 3}},
			expected:     false,
		},
		{
			name:         "shared prerequisite multiple dependents",
			numCourses:   4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 0}},
			expected:     true,
		},
		{
			name:         "long chain then cycle at end",
			numCourses:   5,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}, {2, 4}},
			expected:     false,
		},
		{
			name:         "two separate cycles",
			numCourses:   6,
			prerequisites: [][]int{{1, 0}, {0, 1}, {3, 2}, {2, 3}},
			expected:     false,
		},
		{
			name:         "converging paths no cycle",
			numCourses:   5,
			prerequisites: [][]int{{4, 0}, {4, 1}, {4, 2}, {4, 3}},
			expected:     true,
		},
		{
			name:         "back-edge creates cycle in longer path",
			numCourses:   4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}},
			expected:     false,
		},
		{
			name:         "tree structure no cycle",
			numCourses:   7,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {4, 1}, {5, 2}, {6, 2}},
			expected:     true,
		},
		{
			name:         "single course no prerequisites",
			numCourses:   1,
			prerequisites: [][]int{},
			expected:     true,
		},
		{
			name:         "course requires two which both require a third",
			numCourses:   4,
			prerequisites: [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}},
			expected:     true,
		},
		// Large test cases
		{
			name:       "100 courses linear chain no cycle",
			numCourses: 100,
			// course i requires course i-1, for i in [1..99]
			prerequisites: func() [][]int {
				p := make([][]int, 99)
				for i := 1; i < 100; i++ {
					p[i-1] = []int{i, i - 1}
				}
				return p
			}(),
			expected: true,
		},
		{
			name:       "100 courses linear chain with back-edge cycle",
			numCourses: 100,
			// chain 1→0, 2→1, ..., 99→98 plus 0→99 which creates a cycle
			prerequisites: func() [][]int {
				p := make([][]int, 100)
				for i := 1; i < 100; i++ {
					p[i-1] = []int{i, i - 1}
				}
				p[99] = []int{0, 99}
				return p
			}(),
			expected: false,
		},
		{
			name:       "100 courses binary tree structure no cycle",
			numCourses: 100,
			// course i depends on course (i-1)/2 — binary tree of dependencies
			prerequisites: func() [][]int {
				p := make([][]int, 0, 99)
				for i := 1; i < 100; i++ {
					p = append(p, []int{i, (i - 1) / 2})
				}
				return p
			}(),
			expected: true,
		},
		{
			name:       "100 courses star no cycle",
			numCourses: 100,
			// all courses 1..99 require course 0
			prerequisites: func() [][]int {
				p := make([][]int, 99)
				for i := 1; i < 100; i++ {
					p[i-1] = []int{i, 0}
				}
				return p
			}(),
			expected: true,
		},
		{
			name:       "100 courses with mid-chain cycle",
			numCourses: 100,
			// chain 1→0, 2→1, ..., 99→98 plus a back-edge 50→75 (75 is later in chain, creating cycle)
			prerequisites: func() [][]int {
				p := make([][]int, 100)
				for i := 1; i < 100; i++ {
					p[i-1] = []int{i, i - 1}
				}
				// 50 depends on 75: since 75 already depends on ... depends on 50, this is a cycle
				p[99] = []int{50, 75}
				return p
			}(),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := canFinish(tt.numCourses, tt.prerequisites)
			if actual != tt.expected {
				t.Errorf("canFinish(%d, ...) = %v, want %v", tt.numCourses, actual, tt.expected)
			}
		})
	}
}
