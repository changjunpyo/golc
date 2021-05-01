package problems

import (
	"fmt"
	"testing"
)

type examples88 struct {
	params88
	ans88
}

type params88 struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

type ans88 struct {
}

func Test_Problems88(t *testing.T) {

	ex := []examples88{
		{
			params88{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			ans88{},
		}, {
			params88{[]int{1}, 1, []int{}, 0},
			ans88{},
		},
	}
	fmt.Println("-------------------- Problem 88. Subsets --------------------")
	for _, e := range ex {
		_, p := e.ans88, e.params88

		fmt.Printf("Input: %v %v\n", p.nums1, p.nums2)
		merge(p.nums1, p.m, p.nums2, p.n)
		fmt.Printf("Output: %v\n", p.nums1)
	}
	fmt.Print("\n\n")
}
