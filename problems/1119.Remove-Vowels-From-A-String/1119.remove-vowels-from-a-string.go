package problems

import "strings"

func removeVowels(s string) string {
	var builder strings.Builder

	for _, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			continue
		}
		builder.WriteRune(v)
	}
	return builder.String()
}
