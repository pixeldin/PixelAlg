package linklist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintLink(l *ListNode) {
	for l != nil {
		fmt.Printf("%v -> ", l.Val)
		l = l.Next
	}
	fmt.Println("nil")
}
