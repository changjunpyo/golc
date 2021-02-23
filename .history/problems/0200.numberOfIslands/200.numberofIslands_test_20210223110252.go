package problems

import (
	"fmt"
	"testing"
)

type examples200 struct {
	params200
	ans200
}

type params200 struct {
	grid [][]byte
}

type ans200 struct {
	ans int
}

func Test_Problems200(t *testing.T) {

	ex := []examples200{
		{
			params200{[][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}},
			ans200{1},
		},
	}
	fmt.Println("-------------------- Problem 200.--------------------")
	for _, e := range ex {
		_, p := e.ans200, e.params200
		fmt.Printf("Input: %v\nOutput: %v\n", p.grid,
			numIslands(p.grid))
	}
	fmt.Print("\n\n")
}
