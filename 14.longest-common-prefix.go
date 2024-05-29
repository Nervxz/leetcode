package main

// Longest Common Prefix
//
// # Easy
//
// https://leetcode.com/problems/longest-common-prefix/
//
// Write a function to find the longest common prefix string amongst an array of
// strings.
//
// If there is no common prefix, return an empty string `""`.
//
// **Example 1:**
//
// ```
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
//
// ```
//
// **Example 2:**
//
// ```
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
//
// ```
//
// **Constraints:**
//
// - `1 <= strs.length <= 200`
// - `0 <= strs[i].length <= 200`
// - `strs[i]` consists of only lowercase English letters.
func longestCommonPrefix(strs []string) string {
	// Check case if it is empty
	if len(strs) == 0 {
		return ""
	}
	/// Assign base for the first string
	base := strs[0]
	// Iterate through the strings
	for i := 0; i < len(base); i++ {
		// Iterate through the characters of the strings
		for _, str := range strs[1:] {
			// Check if the index is out of range or the character is not the same
			if i == len(str) || str[i] != base[i] {
				// Return the prefix
				return base[:i]
			}
		}
	}
	return base
}
