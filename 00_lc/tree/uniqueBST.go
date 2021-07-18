package tree

// 96. unique bin tree
// 动态规划, 答案树可以拆解为所有左子树与右子树的组合
/*
	G[0] = G[1] = 1
	G[n] = F[1, n] + F[2, n] + F[i, n]  1<= i <=n

	total = left * right
	F[i,j] = G[i-1] * G[n-i]
	so,
	G[n] = F[1,n]
*/
func numTrees(n int) int {
	ret := make([]int, n+1)
	ret[0] = 1
	ret[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			ret[i] += ret[j-1] * ret[i-j]
		}
	}
	return ret[n]
}
