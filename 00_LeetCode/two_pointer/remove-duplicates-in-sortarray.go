package two_pointer

import "fmt"

func RemoveDuplicates(nums []int)  {
	duplicates := removeDuplicates(nums)
	fmt.Println(duplicates)
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	//def := len(nums)
	for i,j := 0,1; j < len(nums); j++ {
		if nums[j] == nums[i] {
			//def--
			nums = append(nums[:i], nums[i+1:]...)
			//i = j
			j--
		} else {
			i++
		}
	}
	return len(nums)
}