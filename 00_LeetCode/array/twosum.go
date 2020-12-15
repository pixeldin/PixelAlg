package main

import "fmt"

func main()  {
	test := []int{1,9,3,4}
	sum := twoSum(test, 7)
	fmt.Print(sum)
}

func twoSum(nums []int, target int) []int {
	tag := make([]int, 2)
	for i := 0; i < len(nums) - 1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target	{
				tag[0] = i
				tag[1] = j
			}
		}
	}
	return tag
}

func willPass(a, b int) bool {
	return a < b
}

