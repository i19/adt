package booking

// Rotate the string based on the given integer values
// for leftshift and rightshift, all test cases passed.
// abcde 1 -> bcdea
// abcde 2 -> cdeba
// abcde 3 -> decab
// abcde 4 -> edabc
func rotateString(s string, n int) string {
	reverse := func(s1 []rune) {
		for left, right := 0, len(s1)-1; left < right; {
			s1[left], s1[right] = s1[right], s1[left]
			left++
			right--
		}
	}
	x := []rune(s)
	reverse(x)
	reverse(x[:n])
	reverse(x[n:])
	return string(x)
}
