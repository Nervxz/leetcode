package main

// Search Insert Position
//
// Easy
//
// https://leetcode.com/problems/search-insert-position
//
// Given a sorted array of distinct integers and a target value, return the
// index if the target is found. If not, return the index where it would be if
// it were inserted in order.
//
// You must write an algorithm with `O(log n)` runtime complexity.
//
// **Example 1:**
//
// ```
// Input: nums = [1,3,5,6], target = 5
// Output: 2
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [1,3,5,6], target = 2
// Output: 1
//
// ```
//
// **Example 3:**
//
// ```
// Input: nums = [1,3,5,6], target = 7
// Output: 4
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 104`
// - `-104 <= nums[i] <= 104`
// - `nums` contains **distinct** values sorted in **ascending** order.
// - `-104 <= target <= 104`

// Find the insert position of target in a sorted array
func searchInsert(nums []int, target int) int {
	// Initialize the left and right pointers
	l, r := 0, len(nums)-1
	// Binary search
	for l <= r {
		// Calculate the middle index
		mid := (l + r) / 2
		// If the middle element is less than the target, move the left pointer to mid + 1
		if nums[mid] < target {
			l = mid + 1
			// If the middle element is greater than the target, move the right pointer to mid - 1
		} else if nums[mid] > target {
			r = mid - 1
			// If the middle element is equal to the target, return the middle index
		} else {
			return mid
		}
	}
	return l
}
