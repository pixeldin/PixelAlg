package array

func TwoSum(nums []int, target int) []int {
	tag := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			//fmt.Println("i: ", i, "j: ", j)
			if nums[i]+nums[j] == target {
				tag[0] = i
				tag[1] = j
				return tag
			}
		}
	}
	return tag
}

func BetterTwoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		love := target - nums[i]
		if _, ok := m[love]; ok {
			return []int{m[love], i}
		}
		m[nums[i]] = i
	}
	return nil
}
