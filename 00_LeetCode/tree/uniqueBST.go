package tree

func numTrees(n int) int {
	rec := make([]int, 0)
	rec[0] = 0
	rec[1] = 0
	for i := 2; i < n; i++ {

		rec[i] +=
	}
	return rec[n]
}