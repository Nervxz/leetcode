package main

// Same Tree
//
// Easy
//
// https://leetcode.com/problems/same-tree/
//
// Given the roots of two binary trees `p` and `q`, write a function to check if
// they are the same or not.
//
// Two binary trees are considered the same if they are structurally identical,
// and the nodes have the same value.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)
//
// ```
// Input: p = [1,2,3], q = [1,2,3]
// Output: true
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg)
//
// ```
// Input: p = [1,2], q = [1,null,2]
// Output: false
//
// ```
//
// **Example 3:**
//
// ![](https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg)
//
// ```
// Input: p = [1,2,1], q = [1,1,2]
// Output: false
//
// ```
//
// **Constraints:**
//
// - The number of nodes in both trees is in the range `[0, 100]`.
// - `-104 <= Node.val <= 104`
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// isSameTree checks if two binary trees are the same.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Define for a binary tree node
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	// If both nodes are nil, the trees are the same
	if p == nil && q == nil {
		return true
	}

	// If one of the nodes is nil, the trees are not the same
	if p == nil || q == nil {
		return false
	}

	// If the values of the nodes are different, the trees are not the same
	if p.Val != q.Val {
		return false
	}

	// Recursively check the left and right subtrees
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
