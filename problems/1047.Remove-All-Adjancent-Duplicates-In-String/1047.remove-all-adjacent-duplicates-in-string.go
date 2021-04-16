package problems

func removeDuplicates(S string) string {
	s := make([]byte, 0)

	for i:=0; i<len(S); i++{
		if len(s) != 0 && s[len(s)-1] == S[i]{
			s = s[:len(s)-1]
		} else{
			s = append(s, S[i])
		}
	}
	return string(s)
}