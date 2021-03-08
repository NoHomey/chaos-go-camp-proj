package base64

//Test tests if the given string is in base64 encoding.
func Test(s string) bool {
	for _, c := range s {
		if !(isInRange(c, 'A', 'Z') || isInRange(c, 'a', 'z') || c == '-' || c == '_') {
			return false
		}
	}
	return true
}

func isInRange(c, a, b rune) bool {
	return a <= c && c <= b
}
