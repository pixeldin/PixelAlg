package dp

func GetFibSum(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return GetFibSum(n-1) + GetFibSum(n-2)
}

// 509. 剪枝法
func GotFibSumWithMark(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	mark := make(map[int]int, 0)
	mark[0] = 0
	mark[1] = 1
	for i := 2; i <= n; i++ {
		mark[i] = mark[i-1] + mark[i-2]
	}
	return mark[n]
}
