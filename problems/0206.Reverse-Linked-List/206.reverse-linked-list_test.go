package problems

import (
	"fmt"
	"testing"

	"github.com/changjunpyo/golang-leetcode/data-structures"
)

type examples1704 struct {
	params1704
	ans1704
}

type params1704 struct {
	nums1 []int
}

type ans1704 struct {
	ans []int
}

func Test_Problems1704(t *testing.T) {

	ex := []examples1704{
		{
			params1704{[]int{1, 2, 3, 4, 5}},
			ans1704{[]int{5, 4, 3, 2, 1}},
		},
	}
	fmt.Println("-------------------- Problem 1704. Subsets --------------------")
	for _, e := range ex {
		_, p := e.ans1704, e.params1704
		fmt.Printf("Input: %v\nOutput: %v\n", p.nums1,
			data_structures.ListNodeToIntSlice(reverseList(data_structures.IntSliceToListNode(p.nums1))))
	}
	fmt.Print("\n\n")
}
