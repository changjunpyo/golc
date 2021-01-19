package problems

import (
	"fmt"
	"testing"
)

type examples67 struct {
	params67
	ans67
}

type params67 struct {
	str1 string
	str2 string
}

type ans67 struct {
	ans string
}

func Test_Problems67(t *testing.T) {

	ex := []examples67{
		{
			params67{"11","1"},
			ans67{"100"},
		}, {
			params67{"1010", "1011"},
			ans67{"10101"},
		},
	}
	fmt.Println("-------------------- Problem 67. --------------------")
	for _, e := range ex {
		_, p := e.ans67, e.params67
		fmt.Printf("Input: %v %v\nOutput: %v\n", p.str1, p.str2, addBinary(p.str1,p.str2))
	}
	fmt.Print("\n\n")
}
