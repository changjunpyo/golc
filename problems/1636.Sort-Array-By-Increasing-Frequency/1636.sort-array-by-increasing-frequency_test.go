package problems

import (
	"fmt"
	"testing"
)

type examples1636 struct {
	params1636
	ans1636
}

type params1636 struct {
	nums []int
}

type ans1636 struct {
	ans []int
}

func Test_Problems1636(t *testing.T) {

	ex := []examples1636{
		{
			params1636{[]int{1, 1, 2, 2, 2, 3}},
			ans1636{[]int{3, 1, 1, 2, 2, 2}},
		}, {
			params1636{[]int{2, 3, 1, 3, 2}},
			ans1636{[]int{1, 3, 3, 2, 2}},
		}, {
			params1636{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}},
			ans1636{[]int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
		},
	}
	fmt.Println("-------------------- Problem 1636. Subsets --------------------")
	for _, e := range ex {
		_, p := e.ans1636, e.params1636
		fmt.Printf("Input: %v\n Output: %v\n", p.nums, frequencySort(p.nums))
	}
	fmt.Print("\n\n")
}
