package problems

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

type examples122 struct {
	params122
	ans122
}

type params122 struct {
	nums []int
}

type ans122 struct {
	total int
}

func Test_Problems122(t *testing.T) {

	ex := []examples122{
		{
			params122{[]int{7, 1, 5, 3, 6, 4}},
			ans122{7},
		},
		{
			params122{[]int{1, 2, 3, 4, 5}},
			ans122{4},
		},
		{
			params122{[]int{7, 6, 4, 3, 1}},
			ans122{0},
		},
	}
	fmt.Println("-------------------- Problem 122.--------------------")
	for _, e := range ex {
		a, p := e.ans122, e.params122
		got := maxProfit(p.nums)
		want := a.total
		fmt.Printf("Input: %v Output: %v\n", p, got)
		fmt.Printf("got = %d,  want = %d\n", got, want)
		assert.Equal(t, got, want)
	}
	fmt.Print("\n\n")
}
