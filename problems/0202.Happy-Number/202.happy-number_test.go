package problems

import (
	"fmt"
	"testing"
)

type examples202 struct {
	params202
	ans202
}

type params202 struct {
	num int
}

type ans202 struct {
	ans bool
}

func Test_Problems202(t *testing.T) {

	ex := []examples202{
		{
			params202{19},
			ans202{true},
		},
		{
			params202{2},
			ans202{false},
		},
	}
	fmt.Println("-------------------- Problem 202.--------------------")
	for _, e := range ex {
		_, p := e.ans202, e.params202
		fmt.Printf("Input: %v\nOutput: %v\n", p.num,
			isHappy(p.num))
	}
	fmt.Print("\n\n")
}
