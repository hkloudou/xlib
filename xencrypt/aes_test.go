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
	aa := NewAesEnDecrypter(key)
	x2, _ := aa.Encode(d1)
	fmt.Printf("x2:[%d]%2x\n", len(x2), x2)
	x3, _ := aa.Decode(x2)
	fmt.Printf("x3:[%d]%2x\n", len(x3), x3)
	// fmt.Printf("%2x", x2)
	// c := make([]byte, aes.BlockSize+len(d1))
	// iv := c[:aes.BlockSize]
	// x, _ := AesEncrypt(d1, key, key)
	// fmt.Printf("%d,%v\n", len(x), x)
	// fmt.Printf("%2x\n", x)
	// x4, _ := AesDecrypt(x, key, key)
	// fmt.Println("x4", len(x4), x4)
	// fmt.Printf("%2x\n", x4)
	// x3, err := AESEncryptWithNopadding(d1, key, key)
	// fmt.Println(err, len(x3), x3)
	// fmt.Printf("%2x\n", x3)
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
