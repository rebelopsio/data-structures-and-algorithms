package main

// This runs in T = O(N), S = O(N)

var opening = map[rune]bool{
	'{': true,
	'[': true,
	'(': true,
}

var closing = map[rune]bool{
	'}': true,
	']': true,
	')': true,
}

var matching = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

func BalancedBrackets(s string) bool {
	stack := []rune{}
	// Ignore the index here we just want the char
	for _, c := range s {
		// Ignore the value here, we just want the bool val
		if _, found := opening[c]; found {
			stack = append(stack, c)
			continue
		}
		// Ignore the value here, we just want the bool
		if _, found := closing[c]; found {
			// If the current bracket is closing, but there's no opening in the stack, return false
			if len(stack) == 0 {
				return false
				// If the last element in the stack corresponds to the matching element, pop from the stack
			} else if stack[len(stack)-1] == matching[c] {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
	}
	// Ensure that the length of the stack is 0
	return len(stack) == 0
}
