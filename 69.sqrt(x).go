package main

// Sqrt(x)
//
// Easy
//
// https://leetcode.com/problems/sqrtx
//
// Given a non-negative integer `x`, return _the square root of_ `x` _rounded
// down to the nearest integer_. The returned integer should be **non-negative**
// as well.
//
// You **must not use** any built-in exponent function or operator.
//
// - For example, do not use `pow(x, 0.5)` in c++ or `x ** 0.5` in python.
//
// **Example 1:**
//
// ```
// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.
//
// ```
//
// **Example 2:**
//
// ```
// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we round it down
// to the nearest integer, 2 is returned.
//
// ```
//
// **Constraints:**
//
// - `0 <= x <= 231 - 1`

func mySqrt(x int) int {
	// Initialize the left border to 0
	left := 0
	// Initialize the right border to x
	right := x

	// Binary search while the search range is valid
	for left <= right {
		// Calculate the middle of the current search range
		mid := (left + right) / 2
		// If mid squared is less than x, move the left border to mid + 1
		if mid*mid < x {
			left = mid + 1
			// If mid squared is greater than x, move the right border to mid - 1
		} else if mid*mid > x {
			right = mid - 1
			// If mid squared is equal to x, return mid the result
		} else {
			return mid
		}
	}

	// If no exact square root is found, return the right boundary as the result
	return right
}
