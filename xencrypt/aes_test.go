package xencrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"testing"
)

var key = ([]byte("0123456789abcdef"))
var d1 = []byte("0123456789abcdef1")

func Test_aes_1(t *testing.T) {
	// aes.BlockSize
	aa := NewAesEnDecrypter(key, true)
	x2, _ := aa.Encode(d1)
	fmt.Printf("x2:[%d]%2x\n", len(x2), x2)
	x3, _ := aa.Decode(x2)
	fmt.Printf("x3:[%d]%2x %s\n", len(x3), x3, string(x3))
}

func AESEncryptWithNopadding(origData []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// cipher.NewOFB(block, iv).XORKeyStream()
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))

	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}
