package main

// Plus One
//
// Easy
//
// https://leetcode.com/problems/plus-one
//
// You are given a **large integer** represented as an integer array `digits`,
// where each `digits[i]` is the `ith` digit of the integer. The digits are
// ordered from most significant to least significant in left-to-right order.
// The large integer does not contain any leading `0`'s.
//
// Increment the large integer by one and return _the resulting array of
// digits_.
//
// **Example 1:**
//
// ```
// Input: digits = [1,2,3]
// Output: [1,2,4]
// Explanation: The array represents the integer 123.
// Incrementing by one gives 123 + 1 = 124.
// Thus, the result should be [1,2,4].
//
// ```
//
// **Example 2:**
//
// ```
// Input: digits = [4,3,2,1]
// Output: [4,3,2,2]
// Explanation: The array represents the integer 4321.
// Incrementing by one gives 4321 + 1 = 4322.
// Thus, the result should be [4,3,2,2].
//
// ```
//
// **Example 3:**
//
// ```
// Input: digits = [9]
// Output: [1,0]
// Explanation: The array represents the integer 9.
// Incrementing by one gives 9 + 1 = 10.
// Thus, the result should be [1,0].
//
// ```
//
// **Constraints:**
//
// - `1 <= digits.length <= 100`
// - `0 <= digits[i] <= 9`
// - `digits` does not contain any leading `0`'s.

func plusOne(digits []int) []int {
	// start from the last index of the digits array
	index := len(digits) - 1

	// loop to set all trailing 9s to 0
	for index >= 0 && digits[index] == 9 {
		// set the current digit to 0
		digits[index] = 0
		// move to the next digit
		index--
	}
	// if all digits were 9, we need to add a leading 1
	if index == -1 {
		// add a leading 1
		digits = append([]int{1}, digits...)
	} else {
		// else increment the current digit
		digits[index]++
	}

	return digits
}
