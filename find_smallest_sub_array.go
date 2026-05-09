// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
)

/*
Consider two arrays: one representing a paragraph, and one representing a list of keywords. For a given paragraph and keyword list, determine the smallest subarray of the paragraph which covers all of the keywords. The keywords’ order of appearance must match. Keywords may or may not be separated by other words.

Paragraph

"how", "green", "now", "green", "brown", "how", "cow", "now", "brown", "cow"

Keywords

"how", "now", "brown", "cow"

Result

how cow now brown cow

*/

func findSmallestSubArrayBruteForce(paragraph, keywords []string) []string {

	smallest := []string{}
	shortestSubArray := [2]int{0, 0}

	for i := 0; i < len(paragraph); i++ {

		wordIndex := i
		keyWordIndex := 0

		potentialMatchStartIndex, potentialMatchEndIndex := 0, 0

		for {
			if wordIndex > len(paragraph)-1 || keyWordIndex > len(keywords)-1 {

				// if we have found all of the keyword, store this as a potential match
				subArrayLenth := potentialMatchEndIndex - potentialMatchStartIndex
				previousSubArrayLenth := shortestSubArray[1] - shortestSubArray[0]

				// We've found all keywords and the length of the sub array is smaller than any
				// previous sub array we've found
				if keyWordIndex > len(keywords)-1 && ((subArrayLenth < previousSubArrayLenth) || previousSubArrayLenth == 0) {
					shortestSubArray[0] = potentialMatchStartIndex
					shortestSubArray[1] = potentialMatchEndIndex
				}

				// we're at the end of this iteration so we're going to reset our indices and
				// continue looking for sub arrays
				keyWordIndex = 0

				break
			}

			// increment the `keyWordIndex` index whever we find the next
			// word in our keyword list, and also node the start and end
			// indices
			if paragraph[wordIndex] == keywords[keyWordIndex] {

				// Store the indices of potential matches
				if keyWordIndex == 0 {
					potentialMatchStartIndex = wordIndex
				}
				if keyWordIndex == len(keywords)-1 {
					potentialMatchEndIndex = wordIndex
				}

				keyWordIndex++
			}

			wordIndex++
		}
	}

	smallest = paragraph[shortestSubArray[0] : shortestSubArray[1]+1]

	return smallest

}

// This DP implementation trades space for compute. It
// iterates over the paragraph once (although it iterates over the keywords each time, so
// it is O(N*M)). It's space complexity is O(N*M) as well.
func findSmallestSubArrayDP(paragraph, keywords []string) []string {
	smallest := []string{}
	const UNINITIALIZED int = -1
	smallestSubArray := [2]int{UNINITIALIZED, UNINITIALIZED}

	// Build a 2D array with columns being words in the paragraph and rows representing
	// keywords. We will iterate through the paragraph once. When we find any keyword
	// match, we will mark the corresponding cell in the 2D array, with the index where
	// the match occurred.
	// Each value of the array will be populated with the match index, otherwise we will
	// copy the value from the left
	//
	// example:
	// index:  0    1      2    3      4      5    6    7    8      9
	// Word:   how  green  now  green  brown  how  cow  now  brown  cow
	// ----------------------------------------------------------------
	// how     0    0      0    0      0      5    5    5    5      5
	// now     -    -      0    0      0      0    0    5    5      5
	// brown   -    -      -    -      0      0    0    0    5      5
	// cow     -    -      -    -      -      -    0    0    0      5

	// When the final keyword is found, the subset will be: [matrix[row][col], col].
	// For example: for the first 'cow' (index 6), the subset will be [0, 6]. Another
	// subset will be [5,9]

	// build empty array (zero value represents no match)
	matrix := make([][]int, len(keywords))
	for i := 0; i < len(keywords); i++ {
		matrix[i] = make([]int, len(paragraph))
	}

	// iterate over the array and populate the matrix. When a match is detected
	// update the corresponding column of the matrix. When the final keyword is found
	for i, word := range paragraph {
		for k, keyword := range keywords {
			if word == keyword {
				// special case the first row
				if k == 0 {
					matrix[k][i] = i + 1 // adding 1 to all matched values, to account for zero being used to represent no match
				} else {
					matrix[k][i] = matrix[k-1][i-1] // carry over the index where we found the first keyword
				}

				// when we're on the final keyword, identify
				// a valid subarray
				if k == len(keywords)-1 {
					currentSubArrayLength := i - matrix[k][i]
					previousSubArrayLength := smallestSubArray[1] - smallestSubArray[0]

					if currentSubArrayLength < previousSubArrayLength || smallestSubArray[0] == UNINITIALIZED {
						smallestSubArray[0] = matrix[k][i] - 1
						smallestSubArray[1] = i
					}
				}
			} else {
				if i > 0 {
					matrix[k][i] = matrix[k][i-1]
				}
			}
		}
	}

	for _, row := range matrix {
		fmt.Println(row)
	}

	smallest = paragraph[smallestSubArray[0] : smallestSubArray[1]+1]

	return smallest
}
