package main

// Remove Element
//
// Easy
//
// https://leetcode.com/problems/remove-element
//
// Given an integer array `nums` and an integer `val`, remove all occurrences of
// `val` in `nums`
// [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm). The order
// of the elements may be changed. Then return _the number of elements in_
// `nums` _which are not equal to_ `val`.
//
// Consider the number of elements in `nums` which are not equal to `val` be
// `k`, to get accepted, you need to do the following things:
//
// - Change the array `nums` such that the first `k` elements of `nums` contain
// the elements which are not equal to `val`. The remaining elements of `nums`
// are not important as well as the size of `nums`.
// - Return `k`.
//
// **Custom Judge:**
//
// The judge will test your solution with the following code:
//
// ```
// int[] nums = [...]; // Input array
// int val = ...; // Value to remove
// int[] expectedNums = [...]; // The expected answer with correct length.
//                             // It is sorted with no values equaling val.
//
// int k = removeElement(nums, val); // Calls your implementation
//
// assert k == expectedNums.length;
// sort(nums, 0, k); // Sort the first k elements of nums
// for (int i = 0; i < actualLength; i++) {
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
// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// Explanation: Your function should return k = 2, with the first two elements
// of nums being 2.
// It does not matter what you leave beyond the returned k (hence they are
// underscores).
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements
// of nums containing 0, 0, 1, 3, and 4.
// Note that the five elements can be returned in any order.
// It does not matter what you leave beyond the returned k (hence they are
// underscores).
//
// ```
//
// **Constraints:**
//
// - `0 <= nums.length <= 100`
// - `0 <= nums[i] <= 50`
// - `0 <= val <= 100`

func removeElement(nums []int, val int) int {
	// Initialize a counter to 0 to keep track of the number of elements that are not equal to val
	i := 0

	// Iterate through the array of numbers nums using a for loop
	for _, x := range nums {
		// If the current element is not equal to val
		if x != val {
			// place it at the current position of the counter
			nums[i] = x
			// Increment the counter
			i++
		}
	}
	// Return the new length of the array
	return i
}
