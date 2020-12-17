package main

func ThreeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) < 3 {
		return nil
	}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					//fmt.Println(nums[i] , nums [j], nums[k])
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ret
}

func BetterThreeSum(nums []int) [][]int {
	ret := make([][]int, 0)

	return ret
}
