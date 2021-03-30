package _5_binarysearch

/*
	二分查找拓展
*/

// 69.开方
func Sqrt(x int) int {
	if x == 0 {
		return x
	}

	low := 1
	high := x
	for low <= high {
		mid := (low + high) / 2
		sqrt := x / mid
		if sqrt == mid {
			return mid
		} else if sqrt > mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}

// 34. 有序数组 找第一个出现, 最后一个出现的位置
func SearchRange(nums []int, target int) []int {
	ret := make([]int, 0, 2)
	fidx := searchFirstIdx(nums, target)
	lidx := searchLastIdx(nums, target)
	if fidx == -1 && lidx == -1 {
		return []int{-1, -1}
	}
	ret = append(append(ret, fidx), lidx)
	return ret
}

// searchFirstIdx 找第一个出现
func searchFirstIdx(nums []int, target int) int {
	// 123 4 4567
	idx := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		//  找到了, 下标为0 或者 是第1个目标
		if nums[mid] == target && (mid == 0 || (mid != 0 && nums[mid-1] != target)) {
			// 注意越界判断
			return mid
		} else if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return idx
}

// searchLastIdx 找最后一个出现
func searchLastIdx(nums []int, target int) int {
	idx := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		// 找到了, mid是最后1位 或者 mid是最后一个目标
		if nums[mid] == target && (mid == len(nums)-1 || (mid != len(nums)-1 && nums[mid+1] != target)) {
			return mid
		} else if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return idx
}
