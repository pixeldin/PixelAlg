package dp

// 62. 地图走法路径
func UniquePaths(m int, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}
	//var grid [m][n]int
	grid := make([][]int, 0)
	for i := 0; i < m; i++ {
		//for j := 0; j < n; j++ {
		//	line := make([]int, n)
		//}
		line := make([]int, n)
		line[0] = 1
		if i == 0 {
			for j := 1; j < n; j++ {
				line[j] = 1
			}
		}
		grid = append(grid, line)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}
