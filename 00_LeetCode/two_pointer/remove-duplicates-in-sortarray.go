package two_pointer

import "fmt"

func RemoveDuplicates(nums []int) {
	duplicates := removeDuplicates_local(nums)
	fmt.Println(duplicates)
}

/*
	思路: 找到第一个`不重复`的元素覆盖第一个`重复`的元素
*/
func removeDuplicates_local(nums []int) int {
	f := 0
	for i := 1; i < len(nums); i++ {
		if nums[f] != nums[i] {
			f++
			nums[f] = nums[i]
		}
	}
	return f + 1
}
