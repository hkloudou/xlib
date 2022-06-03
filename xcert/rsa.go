package xcert

import (
	"crypto/rand"
	"crypto/rsa"
)

func NewRsa(bits int) *rsa.PrivateKey {
	priv, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil
	}
	return priv
}

func NewRsa256() *rsa.PrivateKey {
	return NewRsa(256)
}
func NewRsa512() *rsa.PrivateKey {
	return NewRsa(512)
}

func NewRsa1024() *rsa.PrivateKey {
	return NewRsa(1024)
}

func NewRsa2048() *rsa.PrivateKey {
	return NewRsa(2048)
}

func NewRsa4096() *rsa.PrivateKey {
	return NewRsa(4096)
}
