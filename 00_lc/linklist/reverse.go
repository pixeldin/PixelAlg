package linklist

/**
* 206.反转链表
  Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// p c h/c n
	//     1   2 3 4 nil
	var pre *ListNode
	c := head
	for c != nil {
		n := c.Next
		c.Next = pre
		pre = c
		c = n
	}
	return pre
}

/**
92. 区间内反转
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 移动指针到右端下标
	//cn := right - left
	return nil
}
