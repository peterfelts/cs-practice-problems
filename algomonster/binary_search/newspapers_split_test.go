package binary_search

import "testing"

func TestNewspapersSplit(t *testing.T) {
	tests := []struct {
		name                string
		newspapersReadTimes []int
		numCoworkers        int
		expected            int
	}{
		{"first example", []int{7, 2, 5, 10, 8}, 2, 18},
		{"second example", []int{2, 3, 5, 7}, 3, 7},
		{"single newspaper", []int{1}, 1, 1},
		{"single maximum read time", []int{100000}, 1, 100000},
		{"single newspaper with extra workers", []int{7}, 100000, 7},
		{"one worker reads everything", []int{7, 2, 5, 10, 8}, 1, 32},
		{"one worker per newspaper", []int{7, 2, 5, 10, 8}, 5, 10},
		{"more workers than newspapers", []int{7, 2, 5, 10, 8}, 8, 10},
		{"two newspapers with one worker", []int{4, 9}, 1, 13},
		{"two newspapers with two workers", []int{4, 9}, 2, 9},
		{"equal times evenly divided", []int{5, 5, 5, 5}, 2, 10},
		{"equal times unevenly divided", []int{5, 5, 5, 5, 5}, 2, 15},
		{"unit times unevenly divided", []int{1, 1, 1, 1, 1, 1, 1}, 3, 3},
		{"exact capacity fills a block", []int{1, 2, 3, 4}, 2, 6},
		{"increasing times", []int{1, 2, 3, 4, 5}, 3, 6},
		{"decreasing times", []int{5, 4, 3, 2, 1}, 3, 6},
		{"largest newspaper first", []int{10, 1, 1, 1}, 2, 10},
		{"largest newspaper last", []int{1, 1, 1, 10}, 2, 10},
		{"largest newspaper in middle", []int{1, 1, 10, 1, 1}, 3, 10},
		{"ordering prevents balanced partition", []int{1, 1, 10, 1, 1}, 2, 12},
		{"nonconsecutive grouping is not allowed", []int{8, 7, 6, 5}, 2, 15},
		{"average time is not a feasible limit", []int{4, 4, 4}, 2, 8},
		{"optimal split near beginning", []int{9, 1, 2, 3}, 2, 9},
		{"optimal split near end", []int{3, 2, 1, 9}, 2, 9},
		{"maximum read times", []int{100000, 100000, 100000}, 2, 200000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := newspapersSplit(tt.newspapersReadTimes, tt.numCoworkers)
			if result != tt.expected {
				t.Errorf("newspapersSplit(%v, %d) = %d, want %d",
					tt.newspapersReadTimes, tt.numCoworkers, result, tt.expected)
			}
		})
	}
}

func TestNewspapersSplitLargeInput(t *testing.T) {
	newspapersReadTimes := make([]int, 100000)
	for i := range newspapersReadTimes {
		newspapersReadTimes[i] = 100000
	}

	tests := []struct {
		name         string
		numCoworkers int
		expected     int64
	}{
		{"one worker with total exceeding 32 bits", 1, 10000000000},
		{"uneven split exceeding 32 bits", 3, 3333400000},
		{"even split", 100, 100000000},
		{"maximum workers", 100000, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := newspapersSplit(newspapersReadTimes, tt.numCoworkers)
			if int64(result) != tt.expected {
				t.Errorf("newspapersSplit(100000 newspapers of 100000 minutes, %d) = %d, want %d",
					tt.numCoworkers, result, tt.expected)
			}
		})
	}
}
