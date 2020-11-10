package _4_sorts

import (
	"fmt"
	"testing"
	"time"
)

func TestCountingSort(t *testing.T) {
	arr := []int{5, 4}
	CountingSort(arr, len(arr))
	t.Log(arr)

	arr = []int{5, 4, 3, 2, 1}
	CountingSort(arr, len(arr))
	t.Log(arr)
}

func TestSleepOrder(t *testing.T)  {
	//runtime.GOMAXPROCS(6)
	target := []int64{0,3,5,1,8,2}
	for _, v := range target {
		go func(n int64) {
			time.Sleep(time.Millisecond * time.Duration(n))
			//time.Sleep(time.Second * time.Duration(n))
			fmt.Println(n)
		}(v)
	}
	for  {
		time.Sleep(1)
	}
}
