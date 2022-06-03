package xcert

import (
	"crypto/rand"
	"math/big"
)

func randomSerialNumber() *big.Int {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, serialNumberLimit)
	return serialNumber
}

// 根据ecdsa密钥生成特征标识码
// func priKeyHash(priv crypto.Signer) []byte {
// 	hash := sha256.New()
// 	switch priv := priv.(type) {
// 	case *ecdsa.PrivateKey:
// 		hash.Write(elliptic.Marshal(priv.Curve, priv.PublicKey.X, priv.PublicKey.Y))
// 		return hash.Sum(nil)
// 	case *rsa.PrivateKey:
// 		hash.Write()
// 	}
// }
