package bit

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	printMaxAndMinWithInt()
	i := reverse(0)
	t.Log(i)
	i = reverse(87654321)
	t.Log(i)

	// max
	i = reverse(7463847412)
	t.Log(i)

	// min
	i = reverse(-8463847412)
	t.Log(i)

	// overflow max
	i = reverse(8463847412)
	t.Log(i)

	// overflow min
	i = reverse(-9463847412)
	t.Log(i)

}

func TestSubsets(t *testing.T) {
	//subsets := Subsets([]int{3, 5, 7})
	//fmt.Println(subsets)

	//SingleNumber([]int{3, 3, 7})
	//fmt.Println(andJudge(3))
	//fmt.Println(hammingWeight(3))
	isPowerOfTwo(3)
	isPowerOfTwo(8)
	isPowerOfTwo(4)
}

// 位运算,判断是否偶数
func andJudge(a int) bool {
	if a&1 == 0 {
		return true
	}
	return false
}

// 192 判断二进制1个数
func hammingWeight(num uint32) int {
	c := 0
	for num > 0 {
		num = num & (num - 1)
		c++
	}
	return c
}

// 231 判断是否是2的pow, 2 [0 1 0], 4 [1,0,0]
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
	//if n == 1 {
	//	return true
	//}
	//c := 1
	//for c < n {
	//	c = c << 1
	//	if c == n {
	//		return true
	//	}
	//}
	//return false
}

// 338. 统计0-num二进制1个数 Counting Bits
func countBits(num int) []int {
	ret := make([]int, num+1)
	//ret[0] = 0
	for i := 1; i <= num; i++ {
		ret[i] = ret[i&(i-1)] + 1
	}
	return ret
}

// 获取num前n位二进制(右往左)
func getFirstNBitFromRight(num, n int) {
	mask := 1<<n - 1
	ret := num & mask
	fmt.Printf("Binary arry: %b\n", ret)
}

func TestFirstNBit(t *testing.T) {
	getFirstNBitFromRight(8, 4)
}
