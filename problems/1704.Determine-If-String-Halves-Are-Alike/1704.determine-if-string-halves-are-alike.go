package problems

func checkVowel(c byte) bool {
	if c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	} else if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	return false
}

func halvesAreAlike(s string) bool {
	count := 0
	for i := 0; i < len(s)/2; i++ {
		if checkVowel(s[i]) {
			count++
		}
		if checkVowel(s[i+len(s)/2]) {
			count--
		}
	}

	if count == 0 {
		return true
	}
	return false
}
