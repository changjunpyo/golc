package problems

import (
	"fmt"
	"testing"
)

type examples70 struct {
	params70
	ans70
}

type params70 struct {
	num int
}

type ans70 struct {
	ans int
}

func Test_Problems70(t *testing.T) {

	ex := []examples70{
		{
			params70{2},
			ans70{2},
		},
		{
			params70{3},
			ans70{3},
		},
	}
	fmt.Println("-------------------- Problem 70.--------------------")
	for _, e := range ex {
		_, p := e.ans70, e.params70
		fmt.Printf("Input: %v\nOutput: %v\n", p.num, climbStairs(p.num))
	}
	fmt.Print("\n\n")
}
