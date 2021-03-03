package secrand

import (
	"crypto/rand"
	"encoding/base64"
)

//RandBytes generates securely random slice of bytes with given length.
func RandBytes(n uint) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//RandString generates securely random string with given length.
func RandString(n uint) (string, error) {
	b, err := RandBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}
