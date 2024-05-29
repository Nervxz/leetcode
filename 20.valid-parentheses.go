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

// Defines a ListNode struct with two fields
// Val (the value of the node) and Next (a pointer to the next node).
type ListNode struct {
	Val  int
	Next *ListNode
}

// Takes two pointers to ListNode (list1 and list2) and returns a pointer to a ListNode.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a new ListNode and assign it to head
	head := &ListNode{}
	current := head
	// Iterate through the lists
	for list1 != nil && list2 != nil {
		// Check if the value of list1 is less than the value of list2
		if list1.Val < list2.Val {
			// Assign the next node to list1
			current.Next = list1
			// Move to the next node
			list1 = list1.Next
		} else {
			// Assign the next node to list2
			current.Next = list2
			// Move to the next node
			list2 = list2.Next
		}
		// Move to the next node
		current = current.Next
	}
	// Check if list1 is not nil
	if list1 != nil {
		// Assign the next node to list1
		current.Next = list1
	} else {
		// Assign the next node to list2
		current.Next = list2
	}
	// Return the next node
	return head.Next
}
