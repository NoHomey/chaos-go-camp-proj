package base64url

import (
	base64 "encoding/base64"
)

//URLTest tests if the given string is in base64 encoding.
func Test(s string) bool {
	for _, c := range s {
		if !(isInRange(c, 'A', 'Z') || isInRange(c, 'a', 'z') || isInRange(c, '0', '9') || c == '-' || c == '_') {
			return false
		}
	}
	return true
}

//Encode encodes given byte slice to string.
func Encode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

//Decode decodes given string.
func Decode(s string) ([]byte, error) {
	bs, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func isInRange(c, a, b rune) bool {
	return a <= c && c <= b
}
