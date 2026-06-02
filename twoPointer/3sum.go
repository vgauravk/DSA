/*
15. 3Sum

Medium

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.


Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

/*
				*********** Algorithm **************

Approach
I think it's tough to manage 3 indices at the same time, so my strategy is to fix i.

Input: nums = [-1,0,1,2,-1,-4]
[-1,0,1,2,-1,-4]
  i
We fix i at index 0 at first. Then we have other two indices j and k. j starts from i + 1 and k starts from the last index.

[-1,0,1,2,-1,-4]
  i j         k
Basically, we calculate a total of 3 numbers, then if the total is 0, that is one of target triplets.

But there are many cases where the total is not 0. In that case we should move j or k to the next, because we fix i at index
0 at first. All calculations in a loop this time, we must use i number.

Question

How can we decide and move j or k to the next index?

My strategy is to sort input array, so that we can decide which index we should move.

[-1,0,1,2,-1,-4]
↓
[-4,-1,-1,0,1,2]
  i  j        k
⭐️ Points

Why do we need to sort input array?
That's because if the total is greater than 0, we want small total next time. In that case, we should move k to the next because
input array is sorted, we are sure that if k move to left index, we will get small total compared with the current total.

On the other hand, if the total is smaller than 0, we want big total next time. In that case we should move j to the right index
to get big total.

Let's see one by one.

[-4,-1,-1,0,1,2]
  i  j        k

nums[i] + nums[j] + nums[k] = -3
We should move j to the next.

I'll speed up

[-4,-1,-1,0,1,2]
  i     j     k

nums[i] + nums[j] + nums[k] = -3

j moves to the next

[-4,-1,-1,0,1,2]
  i       j   k

nums[i] + nums[j] + nums[k] = -2

j moves to the next

[-4,-1,-1,0,1,2]
  i         j k

nums[i] + nums[j] + nums[k] = -1

j moves to the next

[-4,-1,-1,0,1,2]
  i           k
              j
Now j and k are the same index, so we stop iteration. For this question, we have to find all triplets.

Next, we fix i at index 1. j starts from i + 1 and k starts from the last index.

[-4,-1,-1,0,1,2]
     i  j     k
We do the same thing.

[-4,-1,-1,0,1,2]
     i  j     k

nums[i] + nums[j] + nums[k] = 0
We found total 0 and 3 indices are different, so current combination meets condition the description says. That is one of answers.

res = [[-1,-1,2]]
When we found the target triplet, there are several ways to move pointers. I think easiest way is to move j once.

But there is problem. What if the next number is the same as current number like this.

[-4,-1,-1,-1,1,2]
     i    j   k

I changed 0 to -1 at index 3.
In this case, we found [-1,-1,2] again. It's the target triplet but the description says "No duplicate".

To avoid this, when we found one of the target triplets, we move j once but we check the number after we move j. If the number is
the same previous number, we continue to move j until we find differnt number, so that we can avoid duplicate combination.

⭐️ Points

Move j until we find different number when we find one of the target triplets.

Let's go back to main topic.

[-4,-1,-1,0,1,2]
     i    j   k

nums[i] + nums[j] + nums[k] = 1
Now the total is greater than 0. We should move k to the left index to get small total.

[-4,-1,-1,0,1,2]
     i    j k

nums[i] + nums[j] + nums[k] = 0
We found a new triplets that meets the conditions.

res = [[-1,-1,2], [-1,0,1]]
This example is example 1 in the description. It says those two combinations are return value, so I stop explanation.

After that, we do the same thing. j and k are the same index, so we fix i at index 2. j starts from i + 1 and k starts from the
last index.

[-4,-1,-1,0,1,2]
        i j   k
In the end,

return [[-1,-1,2], [-1,0,1]]
*/

package twopointer

import (
	"slices"
)

func ThreeSum(nums []int) [][]int {

	result := [][]int{}

	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {

			total := nums[i] + nums[j] + nums[k]
			if total == 0 {

				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++

				for j < k && nums[j] == nums[j-1] {
					j++
				}

			} else if total > 0 {
				k--
			} else {
				j++
			}

		}
	}

	return result
}
