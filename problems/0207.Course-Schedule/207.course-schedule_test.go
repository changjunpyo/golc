package problems

import (
	"fmt"
	"testing"
)

type examples207 struct {
	params207
	ans207
}

type params207 struct {
	numCourses    int
	prerequisites [][]int
}

type ans207 struct {
	ans bool
}

func Test_Problems207(t *testing.T) {

	ex := []examples207{
		{
			params207{2, [][]int{{1, 0}}},
			ans207{true},
		},
		{
			params207{2, [][]int{{1, 0}, {0, 1}}},
			ans207{false},
		},
	}

	fmt.Println("-------------------- Problem 207.--------------------")
	for _, e := range ex {
		_, p := e.ans207, e.params207
		fmt.Printf("Input: %v %v\nOutput: %v\n", p.numCourses, p.prerequisites,
			canFinish(p.numCourses, p.prerequisites))
	}
	fmt.Print("\n\n")
}
