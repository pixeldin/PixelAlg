package _2_sorts

import (
	"fmt"
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

var arr = []int{2, 7, 6, 5, 4, 1, 3, 0}

func TestQuickSort(t *testing.T) {
	//arr := createRandomArr(10)//[]int{5, 4}
	QuickSort(arr)
	t.Log(arr)

	arr = createRandomArr(5)
	QuickSort(arr)
	t.Log(arr)

}

func TestTopK(t *testing.T) {
	fmt.Println(FindKthLargest(arr, 1))
	fmt.Println(FindKthLargest(arr, 5))
	fmt.Println(FindKthLargest(arr, 8))
}
