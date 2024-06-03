package main

// Climbing Stairs
//
// Easy
//
// https://leetcode.com/problems/climbing-stairs
//
// You are climbing a staircase. It takes `n` steps to reach the top.
//
// Each time you can either climb `1` or `2` steps. In how many distinct ways
// can you climb to the top?
//
// **Example 1:**
//
// ```
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 45`

func climbStairs(n int) int {
	// If n is 1, return 1 because only one way to climb one step
	if n == 1 {
		return 1
	}
	//  Store the number of ways to climb to the previous step
	oneBefore := 1
	// Store the number of ways to climb to two steps before
	twoBefore := 1
	// Store the total number of ways to climb to the current step
	total := 0
	// Loop from step 2 to n
	for i := 2; i <= n; i++ {
		// Calculate the total number of ways to reach the current step
		total = oneBefore + twoBefore
		// Update the total number of ways to climb to two steps before
		twoBefore = oneBefore
		// Update oneBefore to the current total
		oneBefore = total
	}

	return total
}
