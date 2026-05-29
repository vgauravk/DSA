/*
209. Minimum Size Subarray Sum
Medium

Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.



Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1
Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/

package slidingwindow

func MinSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := len(nums) + 1 // Initialize with a value larger than any possible length

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// Shrink the window from the left as long as the sum is >= target
		for sum >= target {
			length := right - left + 1
			if length < minLen {
				minLen = length
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen > len(nums) { // If minLen was never updated, no such subarray exists
		return 0
	}
	return minLen
}
