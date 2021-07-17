package linklist

import (
	"fmt"
	"sync"
	"testing"
)

func buildLink() *ListNode {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	c := &ListNode{3, nil}

	a.Next = b
	b.Next = c
	return a
}

func TestReverse(t *testing.T) {
	ll := buildLink()
	PrintLink(ll)
	rev := reverseList(ll)
	PrintLink(rev)
}

func TestChan(t *testing.T) {
	ch := make(chan int, 1)
	cls := make(chan struct{}, 0)
	//wg := sync.WaitGroup{}
	go func() {
		<-cls
		v, ok := <-ch
		fmt.Printf("val: %v, ch state: %v", v, ok)
	}()
	ch <- 1
	close(ch)
	cls <- struct{}{}

	v, ok := <-ch
	fmt.Printf("val: %v, ch state: %v", v, ok)
}

func TestPanic(t *testing.T) {
	go func() {
		defer func() {
			p := recover()
			fmt.Println("panic happend !", p)
		}()
		panic("haha")
	}()
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}
	//var a,b = []int{},
	//b := copy(a, b)
}

func TestSome(t *testing.T) {
	m := make(map[string]int, 1)
	m[`foo`] = 1

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			m[`foo`]++
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			m[`foo`]++
		}
	}()
	wg.Wait()
}
