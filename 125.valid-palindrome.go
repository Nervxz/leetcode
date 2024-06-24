package main

import (
	"strings"
	"unicode"
)

// Valid Palindrome
//
// Easy
//
// https://leetcode.com/problems/valid-palindrome/
//
// A phrase is a **palindrome** if, after converting all uppercase letters into
// lowercase letters and removing all non-alphanumeric characters, it reads the
// same forward and backward. Alphanumeric characters include letters and
// numbers.
//
// Given a string `s`, return `true` _if it is a **palindrome**, or_ `false`
// _otherwise_.
//
// **Example 1:**
//
// ```
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric
// characters.
// Since an empty string reads the same forward and backward, it is a
// palindrome.
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 2 * 105`
// - `s` consists only of printable ASCII characters.

func isPalindrome(s string) bool {
	// Define the two pointers for the left and right ends of the string
	l, r := 0, len(s)-1

	// Loop until the two pointers meet
	for l < r {
		// If the character at the left pointer is not alphanumeric, move the pointer to the right
		if !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])) {
			l++
			// If the character at the left pointer is not alphanumeric, move the pointer to the right
		} else if !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
			r--
			// If the characters at the left and right pointers are equal, move both pointers towards the center
		} else if strings.ToLower(string(s[l])) == strings.ToLower(string(s[r])) {
			// Move the left pointer to the right
			l++
			r--
		} else {
			// If the characters at the left and right pointers are not equal, the string is not a palindrome
			return false
		}
	}
	// If the loop completes, return palindrome
	return true
}
