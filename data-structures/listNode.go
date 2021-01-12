package data_structures

type ListNode struct {
	Val int
	Next *ListNode
}

// intSliceToListNode convert slice to linkedList
func IntSliceToListNode(nums []int) *ListNode{
	if len(nums) == 0{
		return nil
	}

	l := &ListNode{-1, nil}
	head := l
	for _, v := range nums{
		l.Next = &ListNode{v,nil}
		l = l.Next
	}
	return head.Next
}
// ListNodeToIntSlice convert linkedList to slice
func ListNodeToIntSlice(l *ListNode) []int{
	var res []int

	for l != nil{
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}