package stringutils

// Reverse returns string "in reverse"
func Reverse(s string) string {
	// convert string to tabel of runes
	r := []rune(s)

	// swich first with last etc.
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	// return string
	return string(r)
}
