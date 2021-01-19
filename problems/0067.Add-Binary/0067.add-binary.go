package problems

import "bytes"

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func addBinary(a string, b string) string {
	n, m := len(a)-1, len(b)-1
	if n < m {
		return addBinary(b,a)
	}
	carry := 0
	var sb bytes.Buffer
	for ; n>=0; n,m = n-1,m-1{
		if a[n] == '1'{
			carry++
		}
		if m>=0 && b[m] =='1'{
			carry++
		}

		if carry %2 == 1{
			sb.WriteString("1")
		} else{
			sb.WriteString("0")
		}
		carry /= 2
	}
	if carry == 1{
		sb.WriteString("1")
	}
	return reverse(sb.String())
}