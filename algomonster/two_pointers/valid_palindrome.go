package two_pointers

// Valid Palindrome
// Determine whether a string is a palindrome, ignoring non-alphanumeric characters and case. Examples:

// Input: Do geese see God? Output: True

// Input: Was it a car or a cat I saw? Output: True

// Input: A brown fox jumping over Output: False

func isPalindrome(s string) bool {
	// WRITE YOUR BRILLIANT CODE HERE

	// use two pointers, left and right
	// love them towards each other, checking
	// that the values match on each iteration

	// remove non-alphanumeric characters and change to lowercase
	sProcessed := string("")
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			// shift the rune/ASCII/char value by the char delta between upper
			// and lower case characters (97-65)
			sProcessed += string(s[i] + 'a' - 'A')
		} else if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			sProcessed += string(s[i])
		}
	}

	left, right := 0, len(sProcessed)-1

	for left < right {
		if sProcessed[left] != sProcessed[right] {
			return false
		}
		left++
		right--
	}

	return true
}
