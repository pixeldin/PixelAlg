package _2_sorts

import "fmt"

// 快排拓展
/*
	215. 寻找第k大元素
	思路: 利用快排思想, 每次缩减一半数据量
*/
func FindKthLargest(nums []int, k int) int {
	if k > len(nums) {
		k = len(nums)
	}
	idx := findRank(nums, 0, len(nums)-1, k-1)
	return nums[idx]
}

func findRank(arr []int, l, r, k int) int {
	p := partFork(arr, l, r)
	if p == k {
		return p
	} else if p > k {
		// 偏小, 往左边找 (注意是倒序)
		return findRank(arr, l, p-1, k)
	}
	return findRank(arr, p+1, r, k)
}

func partFork(arr []int, l, r int) int {
	i := l
	point := arr[r]
	for j := l; j < r; j++ {
		// (注意是倒序)
		if arr[j] > point {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}

// 75. 颜色排序, 思路: 计数排序/3路快排
func SortColors(nums []int) {
	i := rankColor(nums, 0, len(nums)-1, 2)
	i = rankColor(nums, 0, len(nums)-1, 1)
	i = rankColor(nums, 0, len(nums)-1, 0)
	fmt.Println(i)
}

func rankColor(arr []int, l, r, k int) int {
	p := partFork(arr, l, r)
	if p == k {
		return p
	} else if p > k {
		// 偏小, 往左边找 (注意是倒序)
		return rankColor(arr, l, p-1, k)
	}
	return rankColor(arr, p+1, r, k)
}
