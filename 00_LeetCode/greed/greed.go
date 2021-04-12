package greed

import (
	"fmt"
)

func JustPrint(arry []int) {
	for i := 0; i < len(arry)-1; i++ {
		for j := i + 1; j < len(arry); j++ {
			fmt.Println(i, " ", j)
		}
	}
	fmt.Println("------------------------")
	for i := 0; i < len(arry); i++ {
		for j := 0; j < len(arry); j++ {
			fmt.Println(i, " ", j)
		}
	}
}

// 322 找零钱问题
func coinChange(coins []int, amount int) int {
	return FindMinCoin(coins, amount)
}

func FindMinCoin(coins []int, amount int) int {
	// 初始化最优结果集
	dpRet := make([]int, amount+1)
	dpRet[0] = 0
	for i := 1; i <= amount; i++ {
		dpRet[i] = -1
	}
	// 自下往上求解
	for i := 1; i <= amount; i++ {
		// 边界最大值
		minCount := amount + 1
		for j := 0; j < len(coins); j++ {
			if coins[j] > i {
				continue
			}
			// 剩下的
			rest := i - coins[j]
			// 尝试拿剩下的最优解
			bestForRest := dpRet[rest]
			if bestForRest == -1 {
				// 无解, 换下一个选择
				continue
			}

			if nc := bestForRest + 1; nc < minCount {
				minCount = nc
			}
		}
		// 记录当前层最优解
		if minCount != (amount + 1) {
			dpRet[i] = minCount
		}
	}

	return dpRet[amount]
}

func coinChangeV2(coins []int, amount int) int {
	// 初始化最优结果集
	dpRet := make([]int, amount+1)
	dpRet[0] = 0
	for i := 1; i <= amount; i++ {
		dpRet[i] = amount + 1
	}
	// 自下往上求解
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] > i {
				continue
			}
			dpRet[i] = min(dpRet[i-coins[j]]+1, dpRet[i])
		}
	}

	if dpRet[amount] > amount {
		return -1
	}

	return dpRet[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
