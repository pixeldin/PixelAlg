package string

import "fmt"

func TestLocal() {
	input := "anviaj"
	ret := lengthOfLongestSubstring(input)
	fmt.Println(ret)

	input = "anbiajbcan"
	ret = lengthOfLongestSubstring(input)
	fmt.Println(ret)
}

func lengthOfLongestSubstring(s string) int {
	unic := make(map[byte]int, 0)
	most := 0
	uniqLen := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := unic[c]; !ok {
			unic[c]++
			uniqLen++
			most = markMost(most, uniqLen)
		} else {
			// mark new one
			unic = make(map[byte]int, 0)
			i = i - uniqLen
			uniqLen = 0
		}
	}
	return most
}

func lengthOfLongestSubstring_todo(s string) int {
	charLocation := make(map[byte]int, 0)
	most := 0
	uniqLen := 0
	var lc byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := charLocation[c]; !ok {
			uniqLen++
		} else {
			if lc != 0 && charLocation[lc] > charLocation[c] {
				uniqLen = i - charLocation[c]
				i = charLocation[c]
			} else if lc != 0 && charLocation[lc] < charLocation[c] {
				uniqLen = i - charLocation[lc]
				i = charLocation[lc]
			}
			lc = c
		}
		charLocation[c] = i
		most = markMost(most, uniqLen)
	}
	return most
}

func markMost(most, uniqLen int) int {
	if uniqLen > most {
		return uniqLen
	}
	return most
}
