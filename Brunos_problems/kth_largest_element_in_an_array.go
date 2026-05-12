package bruno

// Given an integer array nums and an integer k (or n), return the k-th largest element in the array.
// Note: It is the k-th largest element in the sorted order, not the k-th distinct element.

// a := [1, 2, 3, 3, 3, 100, 400, 700]
// [1, 2, 3, 100, 400, 700]
// where k = 3

// kthLargest(a, 1) -> 700
// kthLargest(a, 4) -> 3
// kthLargest(a, 5) -> 3

// k will always be in bounds
// input array is not sorted

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
)

func returnKthLargestBruteForce(values []int, kth int) int {

	sort.Ints(values)

	// a := [1, 2, 3, 3, 3, 100, 400, 700]
	idx := len(values) - kth

	fmt.Println(values)

	return values[idx]

}

func returnKthLargestHeap(values []int, k int) int {

	h := &MaxHeap{}
	heap.Init(h)

	// iterate through the input array and add all items to the heap
	for i := 0; i < len(values); i++ {
		heap.Push(h, values[i])
	}

	fmt.Println(h)

	// As we are told that k will always be in bounds, we simply
	// pop k elements off of the heap and return the kth
	for j := 1; j <= k; j++ {
		max := heap.Pop(h)
		fmt.Println(max)
		if j == k {
			return max.(int)
		}
	}

	return 0
}

func returnKthLargestQuickSelect(values []int, k int) int {

	// Quick Select finds the k-th LOWEST value by default,
	// but we're looking for the k-th LARGEST value,
	// so we adjust k accordingly.
	k = len(values) - k
	return quickSelect(values, 0, len(values)-1, k)
}

// Because we're looking for the kth-largest element, we don't
// need the entire array to be sorted. Taking advantage of this,
// we will use the Quick Select algorithm to find the value we're
// interested in. We don't care if the numbers before or after it
// are sorted, just that it is in the correct kth position.
//
// Quick Select works like this:
//
//  1. Choose a pivot.
//
//  2. Partition every item around that pivot index for a given
//     subset of the array (which will be the entire array during
//     the initial iteration). When we partition a subset of
//     the array, we move every value of the subset to the
//     left or the right of the pivot value. When this process
//     is done, we know that the pivot is in its correct sorted
//     position within the subset, but it won't necessarily be
//     the kth element. After each iteration we can narrow down
//     which side to search next, until the pivot's final position == k.
//
//  3. Compare the pivot index (after partitioning) with k,
//     which gives us three scenarios:
//     a) pivot index == k (we're done!)
//     b) pivot index < k (the value we're looking for is
//     to the right of the pivot index)
//     c) pivot index > k (the value we're looking for is
//     to the left of the pivot index)
//
//  4. Once we know if we're in case b or c, we recurse until
//     the partitioned pivot index == k.
func quickSelect(array []int, left, right, k int) int {

	if left == right {
		return array[left]
	}

	// choose a random pivot point
	randomPivotIndex := left + rand.Intn(right-left)
	fmt.Printf("left: %d, right: %d, randomPivotIndex: %d\n", left, right, randomPivotIndex)

	// Partition the current subarray. This returns the index where
	// the pivot now lives in its correct sorted position.
	newPivot := quickSelectPartition(array, left, right, randomPivotIndex)

	// with this new pivot, determine if we've found what we're looking for (k),
	// or if we need to look to the left or the right
	if newPivot == k {
		return array[k]
	} else if newPivot > k {
		// look left — the pivot landed to the right of k, so k is in the left subarray
		return quickSelect(array, left, newPivot-1, k)
	}
	// look right
	return quickSelect(array, newPivot+1, right, k)
}

func quickSelectPartition(array []int, left, right, pivotIndex int) int {
	// get the value at the pivot index
	pivotValue := array[pivotIndex]

	// keep track of the index that we're considering
	storedIndex := left

	// move the pivot to the end
	swap(array, pivotIndex, right)

	// iterate over the subset, moving all elements that are
	// < pivotValue to the left by swapping them with the stored index
	for i := left; i <= right-1; i++ {
		if array[i] < pivotValue {
			swap(array, i, storedIndex)
			storedIndex++
		}
	}

	// finally, swap the pivot value into its correct sorted position
	swap(array, right, storedIndex)

	return storedIndex

}

func swap(slice []int, a, b int) {
	tmp := slice[b]
	slice[b] = slice[a]
	slice[a] = tmp
}
