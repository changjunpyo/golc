package problems

import (
	"fmt"
	"testing"
)

type examples916 struct {
	params916
	ans916
}

type params916 struct {
	A []string
	B []string
}

type ans916 struct {
	ans []string
}

func Test_Problems916(t *testing.T) {

	ex := []examples916{
		{
			params916{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}},
			ans916{[]string{"facebook", "google", "leetcode"}},
		}, {
			params916{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}},
			ans916{[]string{"apple", "google", "leetcode"}},
		},
		{
			params916{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"}},
			ans916{[]string{"facebook", "google"}},
		},
		{
			params916{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"lo", "eo"}},
			ans916{[]string{"google", "leetcode"}},
		},
	}
	fmt.Println("-------------------- Problem 916. Subsets --------------------")
	for _, e := range ex {
		_, p := e.ans916, e.params916
		fmt.Printf("Input: %v %v\n Output: %v\n", p.A, p.B, wordSubsets(p.A, p.B))
	}
	fmt.Print("\n\n")
}
