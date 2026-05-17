package nickjay

// "123" -> 123
// assume:
// - valid string
// - could have invalid input
// - assum no leading zeros
//
// panic on invalid input

// Add support of octoal or hex

const MODE_BASE10 int = 0
const MODE_HEX int = 1

func stoi(s string, mode int) int {

	m := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
	}

	if len(s) == 0 {
		panic("invalid input")
	}

	// TODO handle sign
	sign := 1
	startIndex := 0
	if s[0] == '-' {
		sign = -1
		startIndex = 1
	}

	runningProduct := 0
	for i := startIndex; i < len(s); i++ {

		// convert each character to int
		cInt := m[s[i]]
		if mode == MODE_BASE10 {
			runningProduct *= 10
		} else {
			runningProduct *= 16
		}
		runningProduct += cInt

	}

	return runningProduct * sign
}
