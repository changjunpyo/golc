package problems

import (
"fmt"
"testing"
)

type examples1047 struct {
	params1047
	ans1047
}

type params1047 struct {
	A string
}

type ans1047 struct {
	ans string
}

func Test_Problems1047(t *testing.T) {

	ex := []examples1047{
		{
			params1047{"abbaca"},
			ans1047{"ca"},
		},
	}
	fmt.Println("-------------------- Problem 1047. Subsets --------------------")
	for _, e := range ex {
		_, p := e.ans1047, e.params1047
		fmt.Printf("Input: %v \nOutput: %v\n", p.A, removeDuplicates(p.A))
	}
	fmt.Print("\n\n")
}
