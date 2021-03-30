package _5_binarysearch

/*
	二分查找
	1234567
	0, 6
	mid := 3
*/
func SearchByDiv(tar int, src []int) int {
	idx := -1
	l, r := 0, len(src)-1
	for l <= r {
		mid := (l + r) / 2
		if src[mid] == tar {
			return mid
		} else if src[mid] > tar {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return idx
}

func BinarySearchV2(src []int, tar, l, r int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if src[mid] == tar {
		return mid
	} else if src[mid] > tar {
		return BinarySearchV2(src, tar, l, mid-1)
	} else {
		return BinarySearchV2(src, tar, mid+1, r)
	}
}
