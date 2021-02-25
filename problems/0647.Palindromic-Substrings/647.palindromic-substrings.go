package problems

func checkPalindromeCount(str string, s, e int) int {
	ret := 0
	if s >= 0 && e < len(str) && str[s] == str[e] {
		ret++
	} else {
		return ret
	}

	if s-1 >= 0 && e+1 < len(str) {
		ret += checkPalindromeCount(str, s-1, e+1)
	}
	return ret
}

func countSubstrings(s string) int {
	ans := 0

	for idx, _ := range s {
		ans += checkPalindromeCount(s, idx, idx)

		ans += checkPalindromeCount(s, idx, idx+1)
	}
	return ans
}
