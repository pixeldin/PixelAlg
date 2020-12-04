package _2_sorts

import (
	"fmt"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	defer PrintTimeCost("mergesort")()
	arr := []int{7, 1, 3, 5}
	MergeSort(arr)
	time.Sleep(200 * time.Millisecond)
	t.Log(arr)
}

func PrintTimeCost(fname string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Cost time: %v(ms), from %s\n", time.Since(start).Milliseconds(), fname)
	}
}
