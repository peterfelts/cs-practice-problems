package main

import (
	"fmt"
)

func validParenthesis(s string) bool {

	// whenever we see a startig parenthesis: {, [, (, we will push the stack
	// when we encounter a closing character: }, ], ), we will pop the stack
	// but when we pop, if we don't end up matching characters, we've found
	// an invalid case

	if len(s) < 2 {
		return false
	}

	// "([][](){{{[]}}})",
	// [ '('
	// [ '(', '[' ]
	//     next character is ']'
	// [ '(' ]
	stack := []rune{}

	for _, v := range s {

		fmt.Println(stack)

		// pop
		if v == ']' || v == '}' || v == ')' {

			// match the corresponding character
			l := stack[len(stack)-1]
			if parentesisCounterpartMatch(l, v) {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		// push
		stack = append(stack, v)

	}

	return len(stack) == 0
}

func parentesisCounterpartMatch(l, r rune) bool {
	return (l == '[' && r == ']') || (l == '{' && r == '}') || (l == '(' && r == ')')
}
