package problems

import "github.com/changjunpyo/golang-leetcode/data-structures"

type ListNode = data_structures.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	var prev *ListNode = nil
	curr := head
	for curr != nil{
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}
