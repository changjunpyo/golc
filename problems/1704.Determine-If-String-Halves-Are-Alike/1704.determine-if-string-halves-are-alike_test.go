package problems

import (
	"fmt"
	"testing"
)

type examples1704 struct {
	params1704
	ans1704
}

type params1704 struct {
	str string
}

type ans1704 struct {
	ans bool
}

func Test_Problems1704(t *testing.T) {

	ex := []examples1704{
		{
			params1704{"book"},
			ans1704{true},
		}, {
			params1704{"textbook"},
			ans1704{false},
		}, {
			params1704{"MerryChristmas"},
			ans1704{false},
		}, {
			params1704{"AbCdEfGh"},
			ans1704{true},
		},
	}
	fmt.Println("-------------------- Problem 1704. Subsets --------------------")
	for _, e := range ex {
		_, p := e.ans1704, e.params1704
		fmt.Printf("Input: %v\nOutput: %v\n", p.str, halvesAreAlike(p.str))
	}
	fmt.Print("\n\n")
}
