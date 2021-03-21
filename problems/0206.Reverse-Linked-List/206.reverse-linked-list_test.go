package problems

import (
	"fmt"
	"testing"

	data_structures "github.com/changjunpyo/golc/dataStructures"
)

type examples206 struct {
	params206
	ans206
}

type params206 struct {
	nums1 []int
}

type ans206 struct {
	ans []int
}

func Test_Problems206(t *testing.T) {

	ex := []examples206{
		{
			params206{[]int{1, 2, 3, 4, 5}},
			ans206{[]int{5, 4, 3, 2, 1}},
		},
	}
	fmt.Println("-------------------- Problem 206.--------------------")
	for _, e := range ex {
		_, p := e.ans206, e.params206
		fmt.Printf("Input: %v\nOutput: %v\n", p.nums1,
			data_structures.ListNodeToIntSlice(reverseList(data_structures.IntSliceToListNode(p.nums1))))
	}
	fmt.Print("\n\n")
}
