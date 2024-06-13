package main

// Valid Parentheses
//
// # Easy
//
// https://leetcode.com/problems/valid-parentheses
//
// Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`,
// `'['` and `']'`, determine if the input string is valid.
//
// An input string is valid if:
//
// 1. Open brackets must be closed by the same type of brackets.
// 2. Open brackets must be closed in the correct order.
// 3. Every close bracket has a corresponding open bracket of the same type.
//
// **Example 1:**
//
// ```
// Input: s = "()"
// Output: true
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "()[]{}"
// Output: true
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "(]"
// Output: false
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 104`
// - `s` consists of parentheses only `'()[]{}'`.

func isValid(s string) bool {
	stack := []rune{} // Initialize an empty stack of runes
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, bracket := range s {
		if _, exists := pairs[bracket]; exists {
			// If the bracket is an opening bracket, push it onto the stack
			stack = append(stack, bracket)
		} else {
			// If the stack is empty or the top of the stack does not match the closing bracket
			if len(stack) == 0 || pairs[stack[len(stack)-1]] != bracket {
				return false
			}
			// Pop the top of the stack
			stack = stack[:len(stack)-1]
		}
	}
	// Return true if the stack is empty, false otherwise
	return len(stack) == 0
}
