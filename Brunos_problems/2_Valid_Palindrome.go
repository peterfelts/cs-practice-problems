package bruno

import (
	"fmt"
	"strings"
)

// 2. Valid Palindrome (Strings & Two Pointers)
// Problem: Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring case.

// Example:

// Input: s = "A man, a plan, a canal: Panama"

// Output: true
func isValidPalindrome(s string) bool {

	fmt.Println(s)

	// remove non alpha-numeric characters
	s = strings.ToLower(s)
	i := 0
	for {
		if i == len(s) || len(s) == 0 {
			break
		}

		// remove non alpha-numeric characters by checkin the ASCII value
		// Note that it would be a lot clearner code to simply use a regex
		// replace of [^a-zA-Z]
		// re := regexp.MustCompile(`[^a-zA-Z]`)
		// result := re.ReplaceAllString(s, "")
		byteValue := byte(s[i])
		if byteValue < byte(48) || /* pre 0-9 */
			(byteValue >= byte(57) && byteValue <= byte(64)) || // : - @
			(byteValue >= byte(91) && byteValue <= byte(96)) || // [ - `]
			byteValue > byte(122) { // z
			s = s[:i] + s[i+1:]
		} else {
			i++
		}
	}

	fmt.Println(s)

	left, right := 0, len(s)-1

	for {
		if left >= right {
			return true
		}

		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
}
