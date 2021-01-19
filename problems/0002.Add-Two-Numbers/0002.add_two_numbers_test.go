package problems

import (
	"fmt"
	"testing"

	"github.com/changjunpyo/golang-leetcode/data-structures"
)

type examples2 struct {
	params2
	ans2
}

type params2 struct {
	nums1 []int
	nums2 []int
}

type ans2 struct {
	ans []int
}

func Test_Problems2(t *testing.T) {

	ex := []examples2{
		{
			params2{[]int{2, 4, 3}, []int{5, 6, 4}},
			ans2{[]int{7, 0, 8}},
		}, {
			params2{[]int{0}, []int{0}},
			ans2{[]int{0}},
		},
		{
			params2{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
			ans2{[]int{8, 9, 9, 9, 0, 0, 0, 1}},
		},
	}
	fmt.Println("-------------------- Problem 2.--------------------")
	for _, e := range ex {
		_, p := e.ans2, e.params2
		fmt.Printf("Input: %v %v\n Output: %v\n", p.nums1, p.nums2,
			data_structures.ListNodeToIntSlice(
				addTwoNumbers(data_structures.IntSliceToListNode(p.nums1),
					data_structures.IntSliceToListNode(p.nums2))))
	}
	fmt.Print("\n\n")
}
