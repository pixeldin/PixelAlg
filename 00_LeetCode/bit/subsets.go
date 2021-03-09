package bit

/*
	求子集
	把子集种类转化为2^len(nums)种组合, 即
*/
func Subsets(nums []int) [][]int {
	total := 1 << len(nums)
	ret := make([][]int, 0)
	for i := 1; i <= total; i++ {
		sing := make([]int, 0)
		tmp := i
		for j := 0; j < len(nums); j++ {
			if tmp&1 == 1 {
				sing = append(sing, nums[j])
			}
			tmp = tmp >> 1
		}
		ret = append(ret, sing)
	}
	return ret
}
