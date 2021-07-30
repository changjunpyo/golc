package problems

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

type examples26 struct {
	params26
	ans26
}

type params26 struct {
	nums []int
}

type ans26 struct {
	k int
	nums []int
}

func Test_Problems26(t *testing.T) {

	ex := []examples26{
		{
			params26{[]int{1,1,2}},
			ans26{2, []int{1,2}},
		},
		{
			params26{[]int{0,0,1,1,1,2,2,3,3,4}},
			ans26{5, []int{0,1,2,3,4}},
		},
	}
	fmt.Println("-------------------- Problem 26.--------------------")
	for _, e := range ex {
		want, p := e.ans26, e.params26
		got := removeDuplicates(p.nums)
		fmt.Printf("Input: %v\nOutput: %v\n", p.nums,
			got)
		assert.Equal(t, got, want.k)

		for i:=0; i< want.k; i++{
			assert.Equal(t, p.nums[i], want.nums[i])
		}
	}
	fmt.Print("\n\n")
}
