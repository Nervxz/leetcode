package main

// Maximum Depth of Binary Tree
//
// Easy
//
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
//
// Given the `root` of a binary tree, return _its maximum depth_.
//
// A binary tree's **maximum depth** is the number of nodes along the longest
// path from the root node down to the farthest leaf node.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)
//
// ```
// Input: root = [3,9,20,null,null,15,7]
// Output: 3
//
// ```
//
// **Example 2:**
//
// ```
// Input: root = [1,null,2]
// Output: 2
//
// ```
//
// **Constraints:**
//
// - The number of nodes in the tree is in the range `[0, 104]`.
// - `-100 <= Node.val <= 100`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	// maxDepth returns the maximum depth of a binary tree
	if root == nil {
		return 0
	}
	// Define for a binary tree node
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	// Recursion to the left and right of the tree
	leftDepth := maxDepth(root.Left)

	// Recursion to the left and right of the tree
	rightDepth := maxDepth(root.Right)

	// Return 1 + the greater of the left and right depth
	return 1 + max(leftDepth, rightDepth)
}
