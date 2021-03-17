package problems

import (
	"fmt"
	"testing"
)

type examples240 struct {
	params240
	ans240
}

type params240 struct {
	nums   [][]int
	target int
}

type ans240 struct {
	ans bool
}

func Test_Problems240(t *testing.T) {

	ex := []examples240{
		{
			params240{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5},
			ans240{true},
		},
	}
	fmt.Println("-------------------- Problem 240. Subsets --------------------")
	for _, e := range ex {
		_, p := e.ans240, e.params240
		fmt.Printf("Input: %v\nOutput: %v\n", p.nums, searchMatrix(p.nums, p.target))
	}
	fmt.Print("\n\n")
}
