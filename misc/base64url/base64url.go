package base64url

import base64 "encoding/base64"

//URLTest tests if the given string is in base64 encoding.
func Test(s string) bool {
	for _, c := range s {
		if !(isInRange(c, 'A', 'Z') || isInRange(c, 'a', 'z') || c == '-' || c == '_') {
			return false
		}
	}
	return true
}

//Encode encodes given byte slice to string.
func Encode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

//DecodeString decodes given string.
func DecodeString(s string) (string, error) {
	d, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func isInRange(c, a, b rune) bool {
	return a <= c && c <= b
}
