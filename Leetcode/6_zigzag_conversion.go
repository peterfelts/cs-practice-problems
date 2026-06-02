package leetcode

// https://leetcode.com/problems/zigzag-conversion/submissions/2020315959/

func convert(s string, numRows int) string {

	if numRows < 2 {
		return s
	}

	// build out array of bytes
	out := make([][]byte, numRows)

	// PAYPALISHIRING
	// P   A   H   N
	// A P L S I I G
	// Y   I   R

	// keep track of direction
	direction := 1 // 1 means down, -1 means up
	row := 0

	// loop over the string and increment row counter, reversing
	// the direction when we get to the last row or to the beginning
	for i := 0; i < len(s); i++ {
		out[row] = append(out[row], s[i])
		row += direction
		if row == numRows-1 {
			direction = -1
		} else if row == 0 {
			direction = 1
		}
	}

	outStr := ""
	for i := 0; i < len(out); i++ {
		outStr += string(out[i])
	}

	return outStr
}
