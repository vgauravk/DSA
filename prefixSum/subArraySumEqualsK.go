/*
560. Subarray Sum Equals K
Medium

Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.



Example 1:

Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2


******************************************************************************
*********************************	Algorithm	******************************
Intuition
A brute force approach would generate all subarrays and compute their sums, leading to O(n²) complexity.

The key observation is:

subarraySum(i...j) = prefixSum[j] - prefixSum[i-1]
We want:

prefixSum[j] - prefixSum[i-1] = k
Rearranging:

prefixSum[i-1] = prefixSum[j] - k
So while traversing the array:

if we've previously seen a prefix sum equal to (currentSum - k),
then a valid subarray ending at current index exists.
We use a hashmap to store:

prefixSum -> frequency
This allows us to count valid subarrays in O(1) time per element.

Approach
Maintain a running prefix sum.
Store frequencies of prefix sums in a hashmap.
For every element:
Update current prefix sum.
Check how many times (sum - k) appeared before.
Add that frequency to answer.
Store/update current prefix sum frequency.
Important:

mp[0] = 1
This represents:

empty prefix before array starts
which helps handle subarrays starting from index 0.


*/

package prefixsum

func SubarraySum(nums []int, k int) int {
	res := 0
	ps := 0

	mp := make(map[int]int)
	mp[0] = 1

	for i := 0; i < len(nums); i++ {
		ps += nums[i]

		if value, isexist := mp[ps-k]; isexist {
			res += value
		}
		mp[ps]++
	}

	return res

}
