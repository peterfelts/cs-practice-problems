/*
	Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

*/

package bruno

// Count the number of zeros. If there is one zero, we can handle that case
// but if there is more than one zero, the product will be zero in all
// cases

// case 1: no zeros -> return the product / nums[i]
// case 2: one zero -> product will be the product of all values other than the one zero
//            -> all values other than when the value of i is zero -> will be zero
//            -> the case where the value of i is zero, we return the product ignoring zero
// case 3: more than one zero -> product will always be zero
func productOfArrayUsingDivision(nums []int) []int {

	// Pass 1: count the number of zeros
	zeroCount := 0
	for i := 0; i < len(nums); i++ {

		// keep track of the number of zeros
		if nums[i] == 0 {
			zeroCount++
		}
	}

	// Pass 2: calculate the product
	// If we have more than one zero, the product will always be zero
	if zeroCount > 1 {
		return make([]int, len(nums))
	}

	// For the case of having one or no zeros, calculate the
	// productZeroIgnored over the entire array (ignoring zero values)
	productZeroIgnored := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			productZeroIgnored *= nums[i]
		}
	}

	// Pass 3: pupulate output array, accounting for the one
	//         zero that we may have
	out := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {

		// one zero (we already accounted for the case where there are > 1 zeros)
		if zeroCount == 1 {
			out[i] = 0

			// when we land on the one non-zero value, the product will be
			// the product of all other numbers, which we calculated already
			if nums[i] == 0 {
				out[i] = productZeroIgnored
			}
		} else {
			// no zeros
			out[i] = productZeroIgnored / nums[i]
		}

	}

	return out
}

// Permutation #2:
//   no use of division
//   O(n) time and space

// i      0  1   2 3
// input [1  2   3 4]
// left  [1* 1   2 6]
// right [24 12  4 1*]
func productOfArrayNoDivision(nums []int) []int {

	// Pass 1: count the number of zeros
	zeroCount := 0
	for i := 0; i < len(nums); i++ {

		// keep track of the number of zeros
		if nums[i] == 0 {
			zeroCount++
		}
	}

	// If we have more than one zero, the product will always be zero
	if zeroCount > 1 {
		return make([]int, len(nums))
	}

	// Pass 2: Populate the left-product and right-product slices

	// store the products to the left and to the right
	// of each index
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	// left
	productL := 1
	productR := 1
	for i := 0; i < len(nums); i++ {
		left[i] = productL
		productL *= nums[i]
		right[len(nums)-1-i] = productR
		productR = productR * nums[len(nums)-1-i]
	}

	// Pass 3: Populate the output array
	out := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		out[i] = left[i] * right[i]
	}

	return out
}
