package dp

// 穷举1
func climbStairs_(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs_(n-1) + climbStairs_(n-2)
}

func ClimbStairs(n int) int {
	var ret = make([]int, 46)
	ret[0] = 0
	ret[1] = 1
	ret[2] = 2
	for i := 3; i <= n; i++ {
		ret[i] = ret[i-1] + ret[i-2]
	}
	return ret[n]
}
