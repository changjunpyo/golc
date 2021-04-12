package problems

import (
	"fmt"
	"testing"
)

type examples1119 struct {
	params1119
	ans1119
}

type params1119 struct {
	str string
}

type ans1119 struct {
	ans string
}

func Test_Problems1119(t *testing.T) {

	ex := []examples1119{
		{
			params1119{"leetcodeisacommunityforcoders"},
			ans1119{"ltcdscmmntyfrcdrs"},
		},
		{
			params1119{"aeiou"},
			ans1119{""},
		},
	}
	fmt.Println("-------------------- Problem 1119.--------------------")
	for _, e := range ex {
		_, p := e.ans1119, e.params1119
		fmt.Printf("Input: %v\nOutput: %v\n", p.str,
			removeVowels(p.str))
	}
	fmt.Print("\n\n")
}
