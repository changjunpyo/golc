package problems

const PrimeRK = 16777619

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	hash, pow := hashStr(needle)
	n := len(needle)

	if needle == haystack[:n] {
		return 0
	}
	h := uint32(0)
	for i := 0; i < n; i++ {
		h = h*PrimeRK + uint32(haystack[i])
	}

	for i := n; i < len(haystack); i++ {
		h *= PrimeRK
		h += uint32(haystack[i])
		h -= pow * (uint32(haystack[i-n]))
		if h == hash && needle == haystack[i-n+1:i+1] {
			return i - n + 1
		}
	}
	return -1

}

func hashStr(str string) (uint32, uint32) {
	h := uint32(0)
	for i := 0; i < len(str); i++ {
		h = h*PrimeRK + uint32(str[i])
	}
	var pow, sq uint32 = 1, PrimeRK
	for i := len(str); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return h, pow
}
