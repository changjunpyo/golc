package problems

import (
	"fmt"
	"testing"
)

type examples48 struct {
	params48
	ans48
}

type params48 struct {
	matrix [][]int
}

type ans48 struct {
	matrix [][]int
}

func Test_Problems48(t *testing.T) {

	ex := []examples48{
		{
			params48{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			ans48{[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		},
	}
	fmt.Println("-------------------- Problem 48.--------------------")
	for _, e := range ex {
		_, p := e.ans48, e.params48
		fmt.Printf("Input: %v\n", p.matrix)
		rotate(p.matrix)
		fmt.Printf("Output: %v\n", p.matrix)
	}
	fmt.Print("\n\n")
}
