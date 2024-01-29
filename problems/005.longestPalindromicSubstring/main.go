package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	begin, maxLen := 0, 1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] && maxLen < 2 {
			begin, maxLen = i, 2
		}
		b, e := i, i

		for e < len(s)-1 && s[e+1] == s[e] {
			e++
		}

		for e < len(s)-1 && b > 0 && s[b-1] == s[e+1] {
			b--
			e++
		}
		if e-b+1 > maxLen {
			begin, maxLen = b, e-b+1
		}
	}

	return s[begin : begin+maxLen]
}
