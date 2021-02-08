package two_pointer

import "fmt"

func MoveZeroes(nums []int) {
	//moveZeroes(nums)
	//moveZeroes([]int{0,1}) // 1,0
	moveZeroes([]int{2, 1}) // 2,1
}

// 将数组中 0 元素都移动到数组的末尾，
// 并且维持所有非 0 元素的相对位置
// int{0,0,1,1,0,2,0,3,3,0})
func moveZeroes(nums []int) {
	f := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			f++
			nums[f] = nums[i]
			if f < i {
				nums[i] = 0
			}
		}
	}
	fmt.Println(nums)
}
