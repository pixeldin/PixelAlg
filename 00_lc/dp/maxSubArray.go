package dp

import "PixelAlg/00_lc/util"

/*
	思路:
	-  如果前面的最大和小于0, 对当前下标来说无益, 则最大收益为我自己:
		if dp[i-1] <= 0, dp[i] = n[i]
	-  如果前面的最大和大于0, 对我来说多多益善, 则我要抱大腿
		if dp[i-1] > 0, dp[i] = n[i] + dp[i-1]
*/
func maxSubArray(nums []int) int {
	// 边界值
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// dp
	dp, res := make([]int, 0), nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = util.Max(dp[i], res)
	}
	return res
}
