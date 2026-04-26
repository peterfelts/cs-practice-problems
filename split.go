// To execute Go code, please declare a func main() in a package "main"

package main

// # String Split

// Without using a built-in string split method, implement a `split(s, c)` method, which receives a string `s` and a character `c` and splits `s` at each occurrence of `c`, returning a list of strings.

// Example 1: s = "split by space", c = ' '
// Output: ["split", "by", "space"]

// Example 2: s = "beekeeper needed", c = 'e'
// Output: ["b", "", "k", "", "p", "r n", "", "d", "d"]

// Example 3: s = "/home/./..//Documents/", c = '/'
// Output: ["", "home", ".", "..", "", "Documents", ""]

// Example 4: s = "", c = '?'
// Output: []

// Constraints:

// - The length of the input string is at most 10^6
// - The delimiter is a single character

func split(s, delimeter string) []string {

	out := []string{}

	if len(s) == 0 {
		return out
	}

	// walk the string, checking for the delimiter as we go.
	// When we find the delimiter, extract the substring and put it
	// into the output array
	i, j := 0, 0
	for i, j = 0, 0; j < len(s); j++ {
		if s[j:j+1] == delimeter {
			sub := s[i:j]
			out = append(out, sub)
			i = j + 1
		}
	}

	// handle the final match
	out = append(out, s[i:j])

	return out
}
