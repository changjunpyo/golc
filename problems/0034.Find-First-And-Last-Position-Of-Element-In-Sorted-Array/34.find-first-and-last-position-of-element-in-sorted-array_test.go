package problems

import (
	"fmt"
	"testing"
)

type examples34 struct {
	params34
	ans34
}

type params34 struct {
	nums   []int
	target int
}

type ans34 struct {
	ans []int
}

func Test_Problems34(t *testing.T) {

	ex := []examples34{
		{
			params34{[]int{5, 7, 7, 8, 8, 10}, 8},
			ans34{[]int{3, 4}},
		}, {
			params34{[]int{5, 7, 7, 8, 8, 10}, 6},
			ans34{[]int{-1, -1}},
		}, {
			params34{[]int{}, 0},
			ans34{[]int{-1, -1}},
		},
	}
	fmt.Println("-------------------- Problem 34.  --------------------")
	for _, e := range ex {
		_, p := e.ans34, e.params34
		fmt.Printf("Input: %v\n Output: %v\n", p.nums, searchRange(p.nums, p.target))
	}
	fmt.Print("\n\n")
}
