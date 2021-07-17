package two_pointer

import "fmt"

func RemoveElement(nums []int, tar int) {
	duplicates := removeElement(nums, tar)
	fmt.Println(duplicates)
}

// 1 2 2 3 2,  2
/*
	思路: 首尾双指针,如果遇到目标元素,用尾元素覆盖,直到
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
			// 只有没有覆盖才移动首指针,因为需要初次比较
			i++
		}
	}
	return l
}
