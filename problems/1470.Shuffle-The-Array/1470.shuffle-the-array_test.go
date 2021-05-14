package problems

import (
	"fmt"
	"testing"
)

type examples1470 struct {
	params1470
	ans1470
}

type params1470 struct {
	nums []int
	n    int
}

type ans1470 struct {
	ans []int
}

func Test_Problems1470(t *testing.T) {

	ex := []examples1470{
		{
			params1470{[]int{2, 5, 1, 3, 4, 7}, 3},
			ans1470{[]int{2, 3, 5, 4, 1, 7}},
		},
		{
			params1470{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4},
			ans1470{[]int{1, 4, 2, 3, 3, 4, 1}},
		},
	}
	fmt.Println("-------------------- Problem 1470.--------------------")
	for _, e := range ex {
		_, p := e.ans1470, e.params1470
		fmt.Printf("Input: %v\nOutput: %v\n", p.nums,
			shuffle(p.nums, p.n))
	}
	fmt.Print("\n\n")
}
