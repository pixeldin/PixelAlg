package linklist

/*
	leetcode 328 迁移偶数位节点到末尾
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := &ListNode{head.Val, nil}
	even := &ListNode{head.Next.Val, nil}
	// mark head
	oh := odd
	eh := even

	i := 0
	cur := head.Next.Next
	for cur != nil {
		if i % 2 == 0 {
			odd.Next = &ListNode{cur.Val, nil}
			odd = odd.Next
		} else {
			even.Next = &ListNode{cur.Val, nil}
			even = even.Next
		}
		i++
		cur = cur.Next
	}
	odd.Next = eh
	return oh
}
