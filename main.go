package main

import (
	slidingwindow "dsa/slidingWindow" // Ensure your go.mod module name is 'dsa'
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

	s := "cbaebabacd"
	p := "abc"
	s = "abab"
	p = "ab"
	res3 := slidingwindow.FindAnagrams(s, p)
	fmt.Println(res3)
}
