package two_pointer

import "fmt"

func RemoveElement(nums []int, tar int) {
	duplicates := removeElement(nums, tar)
	fmt.Println(duplicates)
}

// 1 2 2 3 2,  2
/*
	思路: 首位双指针, 如果遇到目标元素,用尾元素覆盖, 知道
	两指针碰头
*/
func removeElement(nums []int, val int) int {
	l := len(nums)
	for i := 0; i < l; {
		if val == nums[i] {
			//nums[f] = nums[i]
			nums[i] = nums[l-1]
			l--
		} else {
			i++
		}
	}
	return l
}
