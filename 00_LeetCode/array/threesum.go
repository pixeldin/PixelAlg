package main

import (
	"fmt"
	"sort"
)

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
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
				continue
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
				continue
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


func ThreeSumLocal(nums []int) [][]int {
	// 词频map
	m := make(map[int]int, 0)
	// 去重silce
	uiqS := make([]int, 0)
	for _, num := range nums {
		if m[num] == 0 {
			uiqS = append(uiqS, num)
		}
		m[num]++
	}
	// 保证有序
	sort.Ints(uiqS)

	ret := make([][]int, 0)
	// 遍历
	for i := 0; i < len(uiqS); i++ {
		// 0 + 0 + 0
		if uiqS[i] == 0 && m[uiqS[i]] > 3 {
			//fmt.Println("got 0,0,0")
			ret = append(ret, []int{0,0,0})
		}
		for j := i + 1; j < len(uiqS); j++ {
			// 2a+b or 2b+a
			if (2*uiqS[i]+uiqS[j] == 0) && m[uiqS[i]] > 1 {
				fmt.Println("got 2i and j:", uiqS[i], uiqS[i], uiqS[j])
				ret = append(ret, []int{uiqS[i], uiqS[i], uiqS[j]})
				continue
			}
			if (2*uiqS[j]+uiqS[i] == 0) && m[uiqS[j]] > 1 {
				fmt.Println("got 2j and i:", uiqS[j], uiqS[j], uiqS[i])
				ret = append(ret, []int{uiqS[j], uiqS[j], uiqS[i]})
				continue
			}
			if love := 0 - uiqS[i] - uiqS[j]; m[love] > 0 && love > uiqS[j] {
				fmt.Println("got i,j,k:", uiqS[i], uiqS[j], love)
				ret = append(ret, []int{uiqS[i], uiqS[j], love})
			}
		}
	}
	return ret
}
