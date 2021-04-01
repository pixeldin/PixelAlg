package _2_sorts

import (
	"math/rand"
	"testing"
	"time"
)

func createRandomArr(length int) []int {
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(1)
		arr[i] = rand.Intn(10)
	}
	return arr
}

func TestQuickSort(t *testing.T) {
	//arr := createRandomArr(10)//[]int{5, 4}
	arr := []int{2, 7, 6, 5, 4, 1, 3}
	QuickSort(arr)
	t.Log(arr)

	arr = createRandomArr(5)
	QuickSort(arr)
	t.Log(arr)
}
