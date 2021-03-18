package problems

import (
	"fmt"
	"testing"
)

type examples448 struct {
	params448
	ans448
}

type params448 struct {
	nums []int
}

type ans448 struct {
	ans []int
}

func Test_Problems448(t *testing.T) {

	ex := []examples448{
		{
			params448{[]int{4, 3, 2, 7, 8, 2, 3, 1}},
			ans448{[]int{5, 6}},
		},
	}
	fmt.Println("-------------------- Problem 448.  --------------------")
	for _, e := range ex {
		_, p := e.ans448, e.params448
		fmt.Printf("Input: %v\n Output: %v\n", p.nums, findDisappearedNumbers(p.nums))
	}
	fmt.Print("\n\n")
}
