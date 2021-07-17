package two_pointer

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	// 思路
	/*
		第一次遍历:
		求得链表长度l, 最后一个元素位置tail

		k % l 得出最终头结点为倒数第n个
		如果k%l等于0, 返回head

		第二次遍历
		遍历n-1次到达链表断开处,
		让n-1节点指向null作为队尾,
		tail指向head,
		返回n
	 */
	return nil
}
