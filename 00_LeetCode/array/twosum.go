package main

import "fmt"

func main()  {
	test := []int{1,9,3,4,8,10}
	sum := twoSum(test, 18)
	fmt.Print(sum)
}

func twoSum(nums []int, target int) []int {
	tag := make([]int, 2)
	for i := 0; i < len(nums) - 1; i++ {
		for j := i+1; j < len(nums); j++ {
			//fmt.Println("i: ", i, "j: ", j)
			if nums[i] + nums[j] == target	{
				tag[0] = i
				tag[1] = j
				return tag
			}
		}
	}
	return tag
}

func betterTwoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		love := target - nums[i]
		if m[love] == love {
			return []int{m[love], i}
		}
		m[nums[i]] = love
	}
	return nil
}

