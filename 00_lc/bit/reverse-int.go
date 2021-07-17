package bit

import "strconv"

const (
	maxInt = int(^uint32(0) >> 1)
	// 最后一位是7
	//maxInt = 2147483647
	minInt = -maxInt - 1
	// 最后一位是8
	//minInt = -2147483648

	//maxInt64 = int(^uint(0) >> 1) // int64
	//minInt64 = -maxInt64 - 1      // int64
)

func printMaxAndMinWithInt() {
	println(maxInt)
	println(minInt)
	// print reverse num
	max := strconv.Itoa(maxInt)
	println("Reverse Max:")
	for i := len(max); i > 0; i-- {
		c := max[i-1]
		print(string(c))
	}
	println()
	println("Reverse Min:")

	min := strconv.Itoa(minInt)
	for i := len(min); i > 1; i-- {
		c := min[i-1]
		print(string(c))
	}
	println()
}

func reverse(x int) int {
	ret := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if ret > maxInt/10 || (ret == maxInt/10 && pop > 7) {
			return 0
		}
		if ret < minInt/10 || (ret == minInt/10 && pop < -8) {
			return 0
		}
		ret = ret*10 + pop
	}
	return ret
}
