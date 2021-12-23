package xencrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type AesEnDecrypterPadding uint8

const AesEnDecrypterPaddingNO AesEnDecrypterPadding = 0

type AesEnDecrypter struct {
	key []byte
	iv  []byte
	pad AesEnDecrypterPadding
}

func NewAesEnDecrypter(key []byte, iv []byte, pd AesEnDecrypterPadding) *AesEnDecrypter {
	return &AesEnDecrypter{
		key: key,
		iv:  iv,
		pad: pd,
	}
}

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)

	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

//Encode AES 加密
func (m *AesEnDecrypter) Encode(src []byte) (ret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			ret = nil
			err = fmt.Errorf("%v", r)
		}
	}()
	block, err := aes.NewCipher(m.key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, m.key)
	blockmode.CryptBlocks(src, src)
	return src, nil
}

//Decode AES 解密
func (m *AesEnDecrypter) Decode(src []byte) (ret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			ret = nil
			err = fmt.Errorf("%v", r)
		}
	}()
	block, err := aes.NewCipher(m.key)
	if err != nil {
		return nil, err
	}
	blockmode := cipher.NewCBCDecrypter(block, m.key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src, nil
}
