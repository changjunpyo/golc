package problems

import (
	"fmt"
	"testing"
)

type examples647 struct {
	params647
	ans647
}

type params647 struct {
	str string
}

type ans647 struct {
	ans int
}

func Test_Problems647(t *testing.T) {

	ex := []examples647{
		{
			params647{"abc"},
			ans647{3},
		}, {
			params647{"aaa"},
			ans647{6},
		}, {
			params647{"abab"},
			ans647{6},
		},
	}
	fmt.Println("-------------------- Problem 647.--------------------")
	for _, e := range ex {
		_, p := e.ans647, e.params647
		fmt.Printf("Input: %v\nOutput: %v\n", p.str, countSubstrings(p.str))
	}
	fmt.Print("\n\n")
}
