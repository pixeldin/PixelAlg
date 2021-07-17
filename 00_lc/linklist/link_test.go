package linklist

import "testing"

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
