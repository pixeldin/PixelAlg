package main

import "sort"

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
					//sort.Ints(ret)
				}
			}
		}
	}
	return ret
}


/*
	1. 去重并且统计各自频率
	2. 排序
	3. 拆分为 0*3/2a+b/a+2b/abc三情况
 */
func ThreeSumV(nums []int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, value := range nums {
		// every value's count
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	// range of unique
	for i := 0; i < len(uniqNums); i++ {
		// 3 * 0
		if (uniqNums[i] == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] == 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] == 2 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			// u[j] > u[i] !
			love := 0 - (uniqNums[i] + uniqNums[j])
			// ignore 2u[j] + u[i] = 0 (love = u[j])
			if love > uniqNums[j] && counter[love] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], love})
			}
		}
	}
	return res
}
