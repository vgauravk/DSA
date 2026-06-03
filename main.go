package main

import (
	// slidingwindow "dsa/slidingWindow" // Ensure your go.mod module name is 'dsa'
	prefixsum "dsa/prefixSum"
	"fmt"
)

func main() {
	// s := "abcad"
	// res := slidingwindow.LengthOfLongestSubstring(s)
	// fmt.Println(res)

	// target := 4
	// nums := []int{2, 3, 1, 2, 4, 3}
	// res2 := slidingwindow.MinSubArrayLen(target, nums)
	// fmt.Println(res2)

	// s := "cbaebabacd"
	// p := "abc"
	// s = "abab"
	// p = "ab"
	// res3 := slidingwindow.FindAnagrams(s, p)
	// fmt.Println(res3)

	// nums := []int{-1, 0, 1, 2, -1, -4}
	// res4 := twopointer.ThreeSum(nums)
	// fmt.Println(res4)

	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// res5 := twopointer.MaxArea(height)
	// fmt.Println(res5)

	// nums := []int{1, 1, 1}
	// k := 2
	// res6 := prefixsum.SubarraySum(nums, k)
	// fmt.Println(res6)

	nums := []int{1, 1, 1}
	k := 2
	res7 := prefixsum.SubarraySum(nums, k)
	fmt.Println(res7)
}
