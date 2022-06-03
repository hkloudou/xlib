package xcert

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func NewEcdsa(c elliptic.Curve) crypto.Signer {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil
	}
	return priv
}

func NewEcdsaP256() crypto.Signer {
	return NewEcdsa(elliptic.P256())
}

func NewEcdsaP224() crypto.Signer {
	return NewEcdsa(elliptic.P224())
}

func NewEcdsaP384() crypto.Signer {
	return NewEcdsa(elliptic.P384())
}

func NewEcdsaP521() crypto.Signer {
	return NewEcdsa(elliptic.P521())
}
