package xcert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

func randomSerialNumber() *big.Int {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, serialNumberLimit)
	return serialNumber
}

// 根据ecdsa密钥生成特征标识码
func priKeyHash(priKey *ecdsa.PrivateKey) []byte {
	hash := sha256.New()
	hash.Write(elliptic.Marshal(priKey.Curve, priKey.PublicKey.X, priKey.PublicKey.Y))
	return hash.Sum(nil)
}
