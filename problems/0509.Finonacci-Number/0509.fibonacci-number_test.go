package problems

import (
	"fmt"
	"testing"
)

type examples509 struct {
	params509
	ans509
}

type params509 struct {
	num int
}

type ans509 struct {
	ans int
}

func Test_Problems509(t *testing.T) {

	ex := []examples509{
		{
			params509{2},
			ans509{1},
		}, {
			params509{3},
			ans509{4},
		}, {
			params509{20},
			ans509{6765},
		},
	}
	fmt.Println("-------------------- Problem 509. Subsets --------------------")
	for _, e := range ex {
		_, p := e.ans509, e.params509
		fmt.Printf("Input: %v\nOutput: %v\n", p.num, fib(p.num))
	}
	fmt.Print("\n\n")
}
