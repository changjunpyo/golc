package problems

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

type examples217 struct {
	params217
	ans217
}

type params217 struct {
	nums []int
}

type ans217 struct {
	ans bool
}

func Test_Problems217(t *testing.T) {

	ex := []examples217{
		{
			params217{[]int{1,2,3,1}},
			ans217{true},
		},
		{
			params217{[]int{1, 2, 3, 4, 5}},
			ans217{false},
		},
		{
			params217{[]int{1,1,1,3,3,4,3,2,4,2}},
			ans217{true},
		},
	}
	fmt.Println("-------------------- Problem 217.--------------------")
	for _, e := range ex {
		a, p := e.ans217, e.params217
		got := containsDuplicate(p.nums)
		want := a.ans
		fmt.Printf("Input: %v Output: %v\n", p, got)
		fmt.Printf("got = %v,  want = %v\n", got, want)
		assert.Equal(t, got, want)
	}
	fmt.Print("\n\n")
}
