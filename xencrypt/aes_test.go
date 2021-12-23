package xencrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"testing"
)

var key = ([]byte("0123456789abcdef"))
var d1 = []byte("0123456789abcdef")

func Test_aes_1(t *testing.T) {
	// aes.BlockSize

	// x2, _ := NewAesEnDecrypter(key).Encode(d1)
	// t.Log(len(x2), x2)
	// fmt.Printf("%2x", x2)
	// c := make([]byte, aes.BlockSize+len(d1))
	// iv := c[:aes.BlockSize]
	x, _ := AesEncrypt(d1, key, key)
	t.Log(len(x), x)
	fmt.Printf("%2x", x)

	x3, err := AESEncryptWithNopadding(d1, key, key)
	t.Log(err, len(x3), x3)
	fmt.Printf("%2x", x3)
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
