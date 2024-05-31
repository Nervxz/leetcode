package main

// Two Sum
//
// # Easy
//
// https://leetcode.com/problems/two-sum
//
// Given an array of integers `nums` and an integer `target`, return _indices of
// the two numbers such that they add up to `target`_.
//
// You may assume that each input would have **_exactly_ one solution**, and you
// may not use the _same_ element twice.
//
// You can return the answer in any order.
//
// **Example 1:**
//
// ```
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [3,3], target = 6
// Output: [0,1]
//
// ```
//
// **Constraints:**
//
// - `2 <= nums.length <= 104`
// - `-109 <= nums[i] <= 109`
// - `-109 <= target <= 109`
// - **Only one valid answer exists.**
//
// **Follow-up:** Can you come up with an algorithm that is less than
// `O(n2)`time complexity?

// twoSum returns the indices of the two numbers such that they add up to target.
func twoSum(nums []int, target int) []int {
	/*
		Create a map to store the index of the number in the array.
		For example if the array is [2, 7, 11, 15] and the target is 9, the map will be:
		{
			2: 0,
			7: 1,
			11: 2,
			15: 3
		}
	*/
	index := make(map[int]int)

	// Loop over an array of numbers and return the index of the number
	for i, num := range nums {
		// Calculate the difference between the target and the number by subtracting the number from the target
		// For example if the target is 9 and the number is 2, the difference will be 7
		diff := target - num
		// Check if the difference is in the map 
		if j, ok := index[diff]; ok {
			// If the difference is in the map, return the index of the difference and the index of the number
			return []int{j, i}
		}
		// If the difference is not in the map, add the number to the map with the index
		index[num] = i
	}
	// If no solution is found, return nil
	return nil
}
