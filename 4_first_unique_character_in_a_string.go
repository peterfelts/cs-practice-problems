package main

// Problem: Given a string s, find the first non-repeating character in it and return its index.
// If it does not exist, return -1.

// Examples:
// Input: s = "leetcode"
// Output: 0
// Input: s = "loveleetcode"
// Output: 2

type hit = struct {
	character  string
	firstIndex int
	count      int
}

// O(N) time and O(N) space
func firstNonRepeatingCharacter(s string) int {
	// walk the string keep a running acount of each
	// character
	m := make(map[string]hit)

	for i := 0; i < len(s); i++ {
		character := s[i : i+1]
		if item, ok := m[character]; ok {
			item.count++
			m[character] = item

		} else {
			m[character] = hit{character, i, 1}
		}
	}

	// Walk map and keep track of entries we find that
	// have the lowest index and have a count equal to 1
	lowestIndex := len(s)
	for _, v := range m {
		if v.count == 1 && v.firstIndex < lowestIndex {
			lowestIndex = v.firstIndex
		}
	}

	if lowestIndex == len(s) {
		return -1
	}
	return lowestIndex
}
