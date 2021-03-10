package secrand

import "crypto/rand"

//RandBytes generates securely random slice of bytes with given length.
func RandBytes(n uint) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
