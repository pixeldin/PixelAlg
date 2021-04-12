package array

// 242 是否包含相同字母, 可乱序
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mark := make([]int, 26)
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		mark[c]++
	}
	for i := 0; i < len(t); i++ {
		c := t[i] - 'a'
		mark[c]--
	}
	for _, v := range mark {
		if v != 0 {
			return false
		}
	}
	return true
}
