package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var back *ListNode
	// b h n
	// 1 2 3 4
	for head != nil {
		back = head.Next
		tmp := back.Next
		back.Next = head
		head = tmp
	}
	return head
}
