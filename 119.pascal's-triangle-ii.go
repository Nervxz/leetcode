package main

// Pascal's Triangle II
//
// Easy
//
// https://leetcode.com/problems/pascals-triangle-ii/
//
// Given an integer `rowIndex`, return the `rowIndexth` ( **0-indexed**) row of
// the **Pascal's triangle**.
//
// In **Pascal's triangle**, each number is the sum of the two numbers directly
// above it as shown:
//
// ![](https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif)
//
// **Example 1:**
//
// ```
// Input: rowIndex = 3
// Output: [1,3,3,1]
//
// ```
//
// **Example 2:**
//
// ```
// Input: rowIndex = 0
// Output: [1]
//
// ```
//
// **Example 3:**
//
// ```
// Input: rowIndex = 1
// Output: [1,1]
//
// ```
//
// **Constraints:**
//
// - `0 <= rowIndex <= 33`
//
// **Follow up:** Could you optimize your algorithm to use only `O(rowIndex)`
// extra space?

func getRow(rowIndex int) []int {
	// Initialize the row with the first element as 1
	row := make([]int, rowIndex+1)
	row[0] = 1

	// Iterate through the rowIndex
	for i := 1; i <= rowIndex; i++ {
		// Update the row from the end to the beginning
		for j := i; j > 0; j-- {
			// Update the current element by adding the previous element
			row[j] += row[j-1]
		}
	}

	return row
}
