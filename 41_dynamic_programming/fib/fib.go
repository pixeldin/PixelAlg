package fib

// 递归斐波那契
func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

/*
	剪枝: 自下而上把通用结果存起来
 */
func FibWithTmpRet(n int) int {
	ret := make(map[int]int, 0)
	for i := 0; i <= n; i++ {
		if i == 1 || i == 2 {
			ret[i] = 1
			continue
		}
		ret[i] = ret[i - 1] + ret[i - 2]
	}
	return ret[n]
}
