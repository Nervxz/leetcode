package main

// Remove Duplicates from Sorted Array
//
// Easy
//
// https://leetcode.com/problems/remove-duplicates-from-sorted-array
//
// Given an integer array `nums` sorted in **non-decreasing order**, remove the
// duplicates [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm)
// such that each unique element appears only **once**. The **relative order**
// of the elements should be kept the **same**. Then return _the number of
// unique elements in_ `nums`.
//
// Consider the number of unique elements of `nums` to be `k`, to get accepted,
// you need to do the following things:
//
// - Change the array `nums` such that the first `k` elements of `nums` contain
// the unique elements in the order they were present in `nums` initially. The
// remaining elements of `nums` are not important as well as the size of `nums`.
// - Return `k`.
//
// **Custom Judge:**
//
// The judge will test your solution with the following code:
//
// ```
// int[] nums = [...]; // Input array
// int[] expectedNums = [...]; // The expected answer with correct length
//
// int k = removeDuplicates(nums); // Calls your implementation
//
// assert k == expectedNums.length;
// for (int i = 0; i < k; i++) {
//     assert nums[i] == expectedNums[i];
// }
//
// ```
//
// If all assertions pass, then your solution will be **accepted**.
//
// **Example 1:**
//
// ```
// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,_]
// Explanation: Your function should return k = 2, with the first two elements
// of nums being 1 and 2 respectively.
// It does not matter what you leave beyond the returned k (hence they are
// underscores).
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [0,0,1,1,1,2,2,3,3,4]
// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements
// of nums being 0, 1, 2, 3, and 4 respectively.
// It does not matter what you leave beyond the returned k (hence they are
// underscores).
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 3 * 104`
// - `-100 <= nums[i] <= 100`
// - `nums` is sorted in **non-decreasing** order.

// Function to remove duplicates from a slice of integers

func removeDuplicates(nums []int) int {
	// Check if the slice is empty
	if len(nums) == 0 {
		return 0 // If empty, return 0
	}
	// Initialize a counter to 0
	c := 0
	// Loop through the slice starting from the second element
	for i := 1; i < len(nums); i++ {
		// If the current element is not equal to the element at the counter index
		if nums[c] != nums[i] {
			// Increment the counter and update the element at the counter index
			c++
			nums[c] = nums[i]
		}
	}
	// Return the number of unique elements
	return c + 1
}
