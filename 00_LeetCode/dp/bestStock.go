package dp

// 求出最佳股票策略
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	min, total := prices[0], 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if (prices[i] - min) > total {
			total = prices[i] - min
		}
	}
	return total
}
