package main

import "fmt"

var prt = fmt.Printf

func lengthOfLongestSubstring(s string) int {
	for i := 0; i<len(s); i++ {
		sub := s[i]
		for j := i; j<len(s); j++ {
			sub += s[j]
			prt("%s\n", sub)
		}
	}
	return 0
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
