/*
438. Find All Anagrams in a String
Medium

Given two strings s and p, return an array of all the start indices of p's anagrams in s.
You may return the answer in any order.



Example 1:

Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
*/

package slidingwindow

func FindAnagrams(s string, p string) []int {
	result := []int{}

	if len(p) > len(s) {
		return result
	}

	pCount := [26]int{}
	windowCount := [26]int{}

	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
	}

	left := 0

	for right := 0; right < len(s); right++ {
		windowCount[s[right]-'a']++

		if right-left+1 > len(p) {
			windowCount[s[left]-'a']--
			left++
		}

		if windowCount == pCount {
			result = append(result, left)
		}
	}

	return result
}
