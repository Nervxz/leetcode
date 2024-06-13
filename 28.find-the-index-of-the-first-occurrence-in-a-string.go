package main

// Find the Index of the First Occurrence in a String
//
// Easy
//
// https://leetcode.com/find-the-index-of-the-first-occurrence-in-a-string
//
// Given two strings `needle` and `haystack`, return the index of the first
// occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of
// `haystack`.
//
// **Example 1:**
//
// ```
// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.
//
// ```
//
// **Example 2:**
//
// ```
// Input: haystack = "leetcode", needle = "leeto"
// Output: -1
// Explanation: "leeto" did not occur in "leetcode", so we return -1.
//
// ```
//
// **Constraints:**
//
// - `1 <= haystack.length, needle.length <= 104`
// - `haystack` and `needle` consist of only lowercase English characters.

// strStr function to find the first occurrence in a string
func strStr(haystack string, needle string) int {
	// If needle is empty, return 0
	if len(needle) == 0 {
		return 0
	}

	// Create the LPS (Longest Prefix Suffix) array
	lps := make([]int, len(needle))
	pre := 0 // length of the previous longest prefix suffix
	i := 1

	// Preprocess the pattern (needle) to create the LPS array
	for i < len(needle) {
		if needle[i] == needle[pre] {
			pre++        // Increment the length of the previous longest prefix suffix
			lps[i] = pre // Set the LPS value for the current position
			i++          // Move to the next character in the needle
		} else {
			if pre != 0 {
				pre = lps[pre-1] // Update pre to the previous LPS value
			} else {
				lps[i] = 0 // No proper prefix which is also suffix
				i++        // Move to the next character in the needle
			}
		}
	}

	// Start matching the needle with the haystack
	n := 0 // index for needle
	h := 0 // index for haystack

	for h < len(haystack) {
		if needle[n] == haystack[h] {
			n++ // Move to the next character in the needle
			h++ // Move to the next character in the haystack
		}

		if n == len(needle) {
			return h - n // found the needle, return the starting index
		} else if h < len(haystack) && needle[n] != haystack[h] {
			if n != 0 {
				n = lps[n-1] // Update n to the previous LPS value
			} else {
				h++ // Move to the next character in the haystack
			}
		}
	}

	return -1 // needle not found in haystack
}
