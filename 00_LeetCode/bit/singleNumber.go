package bit

func SingleNumber(nums []int) int {
	var tag int
	tag = nums[0]
	for i := 1; i < len(nums); i++ {
		tag = tag ^ nums[i]
	}
	return tag
}
