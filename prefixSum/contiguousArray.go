/*
525. Contiguous Array
Medium

Given a binary array nums, return the maximum length of a contiguous subarray with
an equal number of 0 and 1.



Example 1:

Input: nums = [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
Example 2:

Input: nums = [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
Example 3:

Input: nums = [0,1,1,1,1,1,0,0,0]
Output: 6
Explanation: [1,1,1,0,0,0] is the longest contiguous subarray with equal number of 0 and 1.




Conversation with Gemini
Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.







Example 1:



Input: nums = [0,1]

Output: 2

Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.

Example 2:



Input: nums = [0,1,0]

Output: 2

Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

Example 3:



Input: nums = [0,1,1,1,1,1,0,0,0]

Output: 6

Explanation: [1,1,1,0,0,0] is the longest contiguous subarray with equal number of 0 and 1.

To solve this problem efficiently, we can use a Hash Map (Dictionary) combined with a Prefix Sum technique. This allows us to find the maximum length in a single pass, achieving O(n) time complexity.




******************************************************************************
*********************************	Algorithm	******************************
******************************************************************************

Intuition
To find the maximum length of a contiguous subarray with an equal number of 0s and 1s,
we need to keep track of the cumulative counts of zeros and ones encountered so far.
Whenever the count of zeros becomes equal to the count of ones, it indicates a subarray
with an equal number of both. Additionally, we need to handle the case when there's a
subarray with the same difference between counts of zeros and ones, which suggests that
the number of zeros and ones between those two indices is balanced.

Approach
We can use a dictionary (hash table) to store the difference between counts of zeros
and ones encountered at each index. While iterating through the array, we keep track
of the cumulative counts of zeros and ones. If the difference between these counts
becomes zero, it indicates that the subarray from the beginning to the current index
has an equal number of both. We update the maximum length accordingly. Moreover,
we store the index where each difference between counts first occurred in the
dictionary. If the same difference is encountered again, it means there is a subarray
between the current index and the index stored in the dictionary with an equal number
of zeros and ones. We update the maximum length using this information.

Complexity
Time complexity: O(n)

Space complexity: O(n)




*/

package prefixsum

func FindMaxLength(nums []int) int {
	res := 0
	hmap := make(map[int]int)
	hmap[0] = -1
	zero := 0
	one := 0

	for index, value := range nums {

		if value == 0 {
			zero++
		} else {
			one++
		}

		diff := zero - one
		if val, ok := hmap[diff]; ok {
			res = max(res, index-val)
		} else {
			hmap[diff] = index
		}
	}

	return res
}
