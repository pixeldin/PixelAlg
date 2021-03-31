package dp

// 392. s是否为t的子串, 不必连续, abc is subString of a1b2c3
func IsSubsequence1(s string, t string) bool {
	for len(s) > 0 && len(t) > 0 {
		if s[0] == t[0] {
			s = s[1:]
		}
		t = t[1:]
	}
	return len(s) == 0
}
