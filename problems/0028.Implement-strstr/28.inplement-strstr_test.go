package problems

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

type examples28 struct {
	params28
	ans28
}

type params28 struct {
	haystack string
	needle   string
}

type ans28 struct {
	ans int
}

func Test_Problems28(t *testing.T) {

	ex := []examples28{
		{
			params28{"hello", "ll"},
			ans28{2},
		},
		{
			params28{"aaaaaaaaa", "bba"},
			ans28{-1},
		},
		{
			params28{"", ""},
			ans28{0},
		},
	}
	fmt.Println("-------------------- Problem 28.--------------------")
	for _, e := range ex {
		a, p := e.ans28, e.params28
		got := strStr(p.haystack, p.needle)
		want := a.ans
		fmt.Printf("Input: %v Output: %v\n", p, got)
		fmt.Printf("got = %d,  want = %d\n", got, want)
		assert.Equal(t, got, want)
	}
	fmt.Print("\n\n")
}
