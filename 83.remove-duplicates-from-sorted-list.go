package main

// Remove Duplicates from Sorted List
//
// Easy
//
// https://leetcode.com/problems/remove-duplicates-from-sorted-list
//
// Given the `head` of a sorted linked list, _delete all duplicates such that
// each element appears only once_. Return _the linked list **sorted** as well_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/01/04/list1.jpg)
//
// ```
// Input: head = [1,1,2]
// Output: [1,2]
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/01/04/list2.jpg)
//
// ```
// Input: head = [1,1,2,3,3]
// Output: [1,2,3]
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the list is in the range `[0, 300]`.
// - `-100 <= Node.val <= 100`
// - The list is guaranteed to be **sorted** in ascending order.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// deleteDuplicates removes duplicates from a sorted linked list
func deleteDuplicates(head *ListNode) *ListNode {
	// ListNode defines a node in the linked list
	type ListNode struct {
		Val  int
		Next *ListNode
	}
	// Initialize the current node to the head of the list
	current := head

	// Loop until the current node is not nil and the next node is not nil
	for current != nil && current.Next != nil {
		// Check if the current node value is equal to the next node value
		if current.Val == current.Next.Val {
			// Skip the next node by pointing to the node after the next
			current.Next = current.Next.Next
		} else {
			// Move to the next node in the list
			current = current.Next
		}
	}

	// Return the modified list head
	return head
}
