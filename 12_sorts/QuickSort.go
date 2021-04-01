package _2_sorts

// QuickSort is quicksort methods for golang
func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]

	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				// 交换位置
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return i
}

func Qc(arr []int, l, r int) {
	//l, r := 0, len(arr)
	if l >= r {
		return
	}
	// 划分中界
	p := pt(arr, l, r)
	Qc(arr, l, p-1)
	Qc(arr, p+1, r)
	return
}

func pt(arr []int, l, r int) int {
	point := arr[r]
	i := l
	for j := l; j < r; j++ {
		if arr[j] < point {
			if i != j {
				// 交换
				arr[j], arr[i] = arr[i], arr[j]
			}
			i++
		}
	}
	arr[r], arr[i] = arr[i], arr[r]
	return i
}
