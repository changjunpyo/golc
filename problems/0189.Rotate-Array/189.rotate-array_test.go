package problems

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)


type examples189 struct {
	params189
	ans189
}

type params189 struct {
	nums []int
	k int
}

type ans189 struct {
	ans []int
}

func Test_Problems189(t *testing.T) {

	ex := []examples189{
		{
			params189{[]int{1,2,3,4,5,6,7}, 3},
			ans189{[]int{5,6,7,1,2,3,4}},
		},
		{
			params189{[]int{-1,-100,3,99},2 },
			ans189{[]int{3,99,-1,-100}},
		},

	}
	fmt.Println("-------------------- Problem 189.--------------------")
	fmt.Println("-------------------- Solution1.--------------------")
	for _, e := range ex {
		a, p := e.ans189, e.params189
		got := make([]int, len(p.nums))
		copy(got, p.nums)
		Rotate1(got, p.k)
		want := a.ans
		fmt.Printf("Input: %v Output: %v\n", p, got)
		fmt.Printf("got = %d,  want = %d\n", got, want)
		assert.Equal(t, got, want)
	}
	fmt.Print("\n\n")
	fmt.Println("-------------------- Solution2.--------------------")
	for _, e := range ex {
		a, p := e.ans189, e.params189
		got := make([]int, len(p.nums))
		copy(got, p.nums)
		Rotate2(got, p.k)
		want := a.ans
		fmt.Printf("Input: %v Output: %v\n", p, got)
		fmt.Printf("got = %d,  want = %d\n", got, want)
		assert.Equal(t, got, want)
	}
	fmt.Print("\n\n")
	fmt.Println("-------------------- Solution3.--------------------")
	for _, e := range ex {
		a, p := e.ans189, e.params189
		got := make([]int, len(p.nums))
		copy(got, p.nums)
		Rotate3(got, p.k)
		want := a.ans
		fmt.Printf("Input: %v Output: %v\n", p, got)
		fmt.Printf("got = %d,  want = %d\n", got, want)
		assert.Equal(t, got, want)
	}
	fmt.Print("\n\n")
}
