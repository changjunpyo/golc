package problems

import (
	"fmt"
	"testing"
)

type examples1480 struct {
	params1480
	ans1480
}

type params1480 struct {
	nums []int
}

type ans1480 struct {
	ans []int
}

func Test_Problems1480(t *testing.T) {

	ex := []examples1480{
		{
			params1480{[]int{1, 2, 3, 4}},
			ans1480{[]int{1, 3, 6, 10}},
		},
		{
			params1480{[]int{1, 1, 1, 1, 1}},
			ans1480{[]int{1, 2, 3, 4, 5}},
		},
		{
			params1480{[]int{3, 1, 2, 10, 1}},
			ans1480{[]int{3, 4, 6, 16, 17}},
		},
	}
	fmt.Println("-------------------- Problem 1480.--------------------")
	for _, e := range ex {
		_, p := e.ans1480, e.params1480
		fmt.Printf("Input: %v\nOutput: %v\n", p.nums,
			runningSum(p.nums))
	}
	fmt.Print("\n\n")
}
