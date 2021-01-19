package problems

import (
	"fmt"
	"testing"
)

type examples1 struct {
	params1
	ans1
}

type params1 struct {
	nums   []int
	target int
}

type ans1 struct {
	ans []int
}

func Test_Problems1(t *testing.T) {

	ex := []examples1{
		{
			params1{[]int{2, 7, 11, 15}, 9},
			ans1{[]int{0, 1}},
		}, {
			params1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		}, {
			params1{[]int{3, 3}, 6},
			ans1{[]int{0, 1}},
		},
	}
	fmt.Println("-------------------- Problem 1. --------------------")
	for _, e := range ex {
		_, p := e.ans1, e.params1
		fmt.Printf("Input: %v\n Output: %v\n", p.nums, twoSum(p.nums, p.target))
	}
	fmt.Print("\n\n")
}
