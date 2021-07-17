package two_pointer

// WA, todo...
func trap(height []int) int {
	sum := 0
	length := len(height)
	if length <= 2 {
		return 0
	}
	left := height[0]
	right := height[length-1]

	for i, j := left, right; i < j; {
		if height[left] > height[i+1] {
			// 水往低处流
			i++
		} else {
			// 碰壁, 落袋为安
			sum += (height[i] - height[left]) * (i - left)
			left = i
		}

		if height[right] > height[right-1] {
			j--
		} else {
			sum += (height[j] - height[right]) * (right - j)
			right = j
		}
	}

	return sum
}
