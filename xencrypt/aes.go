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
	// wisepad bool
}

func NewAesEnDecrypter(key []byte, iv []byte) *AesEnDecrypter {
	return &AesEnDecrypter{
		key: key,
		iv:  iv,
		// wisepad: wisepad,
	}
}

func (m *AesEnDecrypter) padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)

	// if m.wisepad && padnum == blocksize {
	// 	pad = []byte{}
	// }
	return append(src, pad...)
}

func (m *AesEnDecrypter) unpadding(src []byte, blocksize int) []byte {
	n := len(src)
	//wisepad is notwork between source len = blocksize-1  and len=blocksize
	// if m.wisepad && n == blocksize {
	// 	return src
	// }
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
	src = m.padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, m.iv)
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
	blockmode := cipher.NewCBCDecrypter(block, m.iv)
	blockmode.CryptBlocks(src, src)
	src = m.unpadding(src, block.BlockSize())
	return src, nil
}
