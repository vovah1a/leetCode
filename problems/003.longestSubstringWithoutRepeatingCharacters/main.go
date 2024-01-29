package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	hm := make(map[uint8]int)
	var left, right, result int

	for right < len(s) {
		r := s[right]
		hm[r] += 1
		for hm[r] > 1 {
			l := s[left]
			hm[l] -= 1
			left += 1
		}
		result = max(result, right-left+1)

		right += 1
	}
	return result
}
