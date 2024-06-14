package main

// Add Binary
//
// Easy
//
// https://leetcode.com/problems/add-binary
//
// Given two binary strings `a` and `b`, return _their sum as a binary string_.
//
// **Example 1:**
//
// ```
// Input: a = "11", b = "1"
// Output: "100"
//
// ```
//
// **Example 2:**
//
// ```
// Input: a = "1010", b = "1011"
// Output: "10101"
//
// ```
//
// **Constraints:**
//
// - `1 <= a.length, b.length <= 104`
// - `a` and `b` consist only of `'0'` or `'1'` characters.
// - Each string does not contain leading zeros except for the zero itself.

// addBinary takes two binary strings a and b, and returns sum
func addBinary(a string, b string) string {
	// Initialize result string to store the final binary sum
	result := ""
	// Initialize carry to store the carry value
	carry := 0

	// Pointers for both strings starting from the end
	i, j := len(a)-1, len(b)-1

	// Loop until both strings are processed
	for i >= 0 || j >= 0 || carry > 0 {
		// Declare a variable sum to store the sum of the current bits and the carry
		var sum int

		// Add the bits from both strings if available
		if i >= 0 {
			// Convert the character at i in string a to an integer and add to sum
			sum += int(a[i] - '0')
			// Decrement the pointer i.
			i--
		}
		if j >= 0 {
			// Convert the character at j in string b to an integer and add to sum
			sum += int(b[j] - '0')
			// Decrement the pointer j
			j--
		}

		// Add the carry to the sum
		sum += carry

		// Compute the new carry. If sum is 2 or more, carry will be 1, otherwise it will be 0
		carry = sum / 2

		// Compute the current bit to be added to the result by taking sum % 2 (which will be 0 or 1)
		// Convert it to a character and prepend it to the result string
		result = string(sum%2+'0') + result
	}

	// Return sum
	return result
}
