package main

// Palindrome Number
//
// # Easy
//
// https://leetcode.com/problems/palindrome-number
//
// Given an integer `x`, return `true` _if_ `x` _is a_ _**palindrome**_ _, and_
// `false` _otherwise_.
//
// **Example 1:**
//
// ```
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
//
// ```
//
// **Example 2:**
//
// ```
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it
// becomes 121-. Therefore it is not a palindrome.
//
// ```
//
// **Example 3:**
//
// ```
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
// ```
//
// **Constraints:**
//
// - `-231 <= x <= 231 - 1`
//
// **Follow up:** Could you solve it without converting the integer to a string?

func isPalindrome(x int) bool {
	// Check negative number
	// Check if the last digit is 0 and the number is not 0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// Reverse half of the number
	reversed := 0
	// When x is less than reversed
	// it means we've reached the middle of the number
	for x > reversed {
		reversed = (reversed * 10) + (x % 10)
		// Remove the last digit from x
		x /= 10
	}
	// Check if the number is a palindrome
	return x == reversed || x == reversed/10
}
