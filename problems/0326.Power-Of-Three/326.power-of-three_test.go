package problems

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

type examples326 struct {
	params326
	ans326
}

type params326 struct {
	num int
}

type ans326 struct {
	ans bool
}

func Test_Problems326(t *testing.T) {

	ex := []examples326{
		{
			params326{27},
			ans326{true},
		},
		{
			params326{0},
			ans326{false},
		},
		{
			params326{9},
			ans326{true},
		},
	}
	fmt.Println("-------------------- Problem 326.--------------------")
	for _, e := range ex {
		a, p := e.ans326, e.params326
		got := isPowerOfThree(p.num)
		want := a.ans
		fmt.Printf("Input: %v Output: %v\n", p, got)
		fmt.Printf("got = %v,  want = %v\n", got, want)
		assert.Equal(t, got, want)
	}
	fmt.Print("\n\n")
}
