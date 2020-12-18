package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	test2 := []int{0, 2, -1, 1, 0, 1, -1, -3, 3, -2}
	sum2 := threeSumLocal(test2)
	fmt.Print(sum2)
}

func threeSumLocal(nums []int) [][]int {
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
		for j := i + 1; j < len(uiqS); j++ {
			// 0 + 0 + 0
			if uiqS[i] == 0 && m[uiqS[i]] == 3 {
				fmt.Println("got 0,0,0")
				ret = append(ret, []int{0, 0, 0})
			}
			// 2a+b or 2b+a
			if (2*uiqS[i]+uiqS[j] == 0) && m[uiqS[i]] > 1 {
				fmt.Println("got 2i and j:", uiqS[i], uiqS[i], uiqS[j])
				ret = append(ret, []int{uiqS[i], uiqS[i], uiqS[j]})
			}
			if (2*uiqS[j]+uiqS[i] == 0) && m[uiqS[j]] > 1 {
				fmt.Println("got 2j and i:", uiqS[j], uiqS[j], uiqS[i])
				ret = append(ret, []int{uiqS[j], uiqS[j], uiqS[i]})
			}
			if love := 0 - uiqS[i] - uiqS[j]; m[love] > 1 && love > uiqS[j] {
				fmt.Println("got i,j,k:", uiqS[i], uiqS[j], love)
				ret = append(ret, []int{uiqS[i], uiqS[j], love})
			}
		}
	}
	return ret
}
