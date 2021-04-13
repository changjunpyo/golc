package problems

import (
"fmt"
"testing"
)

type examples279 struct {
	params279
	ans279
}

type params279 struct {
	num int
}

type ans279 struct {
	ans int
}

func Test_Problems279(t *testing.T) {

	ex := []examples279{
		{
			params279{12},
			ans279{3},
		},
		{
			params279{13},
			ans279{2},
		},
	}
	fmt.Println("-------------------- Problem 279.--------------------")
	for _, e := range ex {
		_, p := e.ans279, e.params279
		fmt.Printf("Input: %v\nOutput: %v\n", p.num,
			numSquares(p.num))
	}
	fmt.Print("\n\n")
}
