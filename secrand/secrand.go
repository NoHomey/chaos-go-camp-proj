package secrand

import (
	"crypto/rand"

	"github.com/NoHomey/chaos-go-camp-proj/misc/base64url"
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
	return base64url.Encode(b), err
}
