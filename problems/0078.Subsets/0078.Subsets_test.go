package problems

import (
	"fmt"
	"testing"
)

type examples78 struct {
	params78
	ans78
}

type params78 struct {
	nums []int
}

type ans78 struct {
	ans [][]int
}

func Test_Problems78(t *testing.T) {

	ex := []examples78{
		{
			params78{[]int{1, 2, 3}},
			ans78{[][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}},
		}, {
			params78{[]int{0}},
			ans78{[][]int{{}, {1}}},
		},
	}
	fmt.Println("-------------------- Problem 78. Subsets --------------------")
	for _, e := range ex {
		_, p := e.ans78, e.params78
		fmt.Printf("Input: %v\n Output: %v\n", p.nums, subsets(p.nums))
	}
	fmt.Print("\n\n")
}
