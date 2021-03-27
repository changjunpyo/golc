package problems

func wordSubsets(A []string, B []string) []string {

	maxB := make([]int, 26)

	for _, val := range B {
		counter := make([]int, 26)
		for i := 0; i < len(val); i++ {
			idx := val[i] - 'a'
			counter[idx]++
			if counter[idx] > maxB[idx] {
				maxB[idx] = counter[idx]
			}
		}
	}

	ans := make([]string, 0)

	for _, val := range A {
		x := count(val)
		check := true
		for i := 0; i < 26; i++ {
			if x[i] < maxB[i] {
				check = false
				break
			}
		}
		if check {
			ans = append(ans, val)
		}
	}

	return ans
}

func count(str string) []int {
	ans := make([]int, 26)
	for i := 0; i < len(str); i++ {
		ans[str[i]-'a']++
	}
	return ans
}
