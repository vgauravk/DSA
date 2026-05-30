/*
76. Minimum Window Substring
Hard

Given two strings s and t of lengths m and n respectively, return the minimum window substring
of s such that every character in t (including duplicates) is included in the window.
If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.



Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.

*/

Step 1: Initialization
Create a frequency map (freq) of all characters in string t.

Initialize two pointers: left = 0 and right = 0.

Set a variable required = len(t). This tracks how many total character matches we still need to complete the window.

Set up two tracking variables for the result: min_len = infinity and start_idx = 0.



Step 2: Expand the Window (Outer Loop)
Iterate right from 0 to len(s) - 1:

Get the character entering the window: ch = s[right].

Check if it's needed: If freq[ch] > 0, it means this character is actively needed to fulfill 
string t. Decrement required by 1.

Update the map: Decrement freq[ch] by 1. (Note: If ch wasn't in t, its frequency becomes 
negative, which is correct).



Step 3: Shrink the Window (Inner Loop)
While required == 0 (meaning the current window [left, right] contains all characters of t):

Update the Answer: Calculate the current window length: current_len = right - left + 1. If current_len < min_len, 
update min_len = current_len and start_idx = left.

Prepare to remove the leftmost character: Get left_ch = s[left].

Update the map: Increment freq[left_ch] by 1.

Check if validity breaks: If freq[left_ch] > 0, it means we just removed a character that is strictly necessary for t, 
and we no longer have enough copies of it in our window. Increment required by 1 (this will break the while loop).

Move the left pointer forward (left++) to try a smaller window.



Step 4: Return the Result
After the outer loop finishes:

If min_len was updated from infinity, return the substring s[start_idx : start_idx + min_len].

Otherwise, return an empty string "".

The Cheat-Sheet Core Logic
The entire algorithm hinges on these two exact conditional checks:



GO
# When RIGHT expands:
if freq[ch] > 0 {
    required--
}
freq[ch]--

# When LEFT shrinks:
freq[left_ch]++
if freq[left_ch] > 0 {
    required++
}