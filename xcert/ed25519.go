package xcert

import (
	"crypto/ed25519"
	"crypto/rand"
)

func NewEd25519() ed25519.PrivateKey {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil
	}
	return priv
}
