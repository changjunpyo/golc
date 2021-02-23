package problems

import (
	"github.com/changjunpyo/golang-leetcode/dataStructures"
)

// Definition for singly-linked list.
type ListNode = dataStructures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	node := head
	carry, total := 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}
		total += carry
		carry = total / 10
		node.Next = &ListNode{Val: total % 10}
		total = 0
		node = node.Next
	}
	if carry > 0 {
		node.Next = &ListNode{Val: 1}
	}
	return head.Next
}
