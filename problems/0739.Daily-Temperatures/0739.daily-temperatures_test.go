package problems

import (
	"fmt"
	"testing"
)

type examples739 struct {
	params739
	ans739
}

type params739 struct {
	nums []int
}

type ans739 struct {
	ans []int
}

func Test_Problems739(t *testing.T) {

	ex := []examples739{
		{
			params739{[]int{73, 74, 75, 71, 69, 72, 76, 73}},
			ans739{[]int{1, 1, 4, 2, 1, 1, 0, 0}},
		},
	}
	fmt.Println("-------------------- Problem 739.--------------------")
	for _, e := range ex {
		_, p := e.ans739, e.params739
		fmt.Printf("Input: %v\nOutput: %v\n", p.nums,
			dailyTemperatures(p.nums))
	}
	fmt.Print("\n\n")
}
