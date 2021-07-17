package array

import (
	"fmt"
	"sort"
)

func ThreeSumClosestLocal() {

	test1 := []int{0, 0, 0, 0}
	ret := threeSumClosest(test1, 0)
	fmt.Println(ret)

	test2 := []int{0, -1, 2, 8, 3, -3, 2}
	ret2 := threeSumClosest(test2, 17)
	fmt.Println(ret2)

	test2 = []int{305, -9, 2, 7, 3, -3, -101}
	ret2 = threeSumClosest(test2, 203)
	fmt.Println(ret2)
}

/*
	Input: nums = [-1,2,1,-4], target = 1
	Output: 2
	Explanation: The sum that is closest to the
	target is 2. (-1 + 2 + 1 = 2).
*/
func threeSumClosest(nums []int, target int) int {
	min := 10001
	sort.Ints(nums)
	ret := nums[len(nums)-1]
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			}
			ret, min = takeLess(min, ret, target, sum)
			if sum < target {
				left++
				continue
			}
			if sum > target {
				right--
				continue
			}
		}
	}
	return ret
}

func intAbs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func takeLess(min, ret, target, sum int) (int, int) {
	if intAbs(target-sum) <= min {
		ret = sum
		return ret, intAbs(target - sum)
	} else {
		return ret, min
	}
}
