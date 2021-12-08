package xencrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type RsaEncrypter struct {
	pub *rsa.PublicKey
	pri *rsa.PrivateKey
}

func NewRsaEncrypter(pri *rsa.PrivateKey, pub *rsa.PublicKey) *RsaEncrypter {
	return &RsaEncrypter{
		pri: pri,
		pub: pub,
	}
}

func NewRsaEncrypterWithPrivateBytesHard(pri []byte) *RsaEncrypter {
	x, err := NewRsaEncrypterWithPrivateBytes(pri)
	if err != nil {
		panic(err)
	}
	return x
}

func NewRsaEncrypterWithPublicBytesHard(pub []byte) *RsaEncrypter {
	x, err := NewRsaEncrypterWithPublicBytes(pub)
	if err != nil {
		panic(err)
	}
	return x
}

func NewRsaEncrypterWithPrivateBytes(pri []byte) (*RsaEncrypter, error) {
	block, _ := pem.Decode(pri)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	return &RsaEncrypter{
		pri: priv,
		pub: &priv.PublicKey,
	}, nil
}

func NewRsaEncrypterWithPublicBytes(pub []byte) (*RsaEncrypter, error) {
	block, _ := pem.Decode(pub)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return &RsaEncrypter{
		pub: pubInterface.(*rsa.PublicKey),
	}, nil
}

// EncodeOAEP 可变长的数据加密
func (m *RsaEncrypter) EncodeOAEP(origData []byte) ([]byte, error) {
	hash := sha256.New()
	msgLen := len(origData)
	step := m.pub.Size() - 2*hash.Size() - 2
	var encryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		encryptedBlockBytes, err := rsa.EncryptOAEP(hash, rand.Reader, m.pub, origData[start:finish], nil)
		if err != nil {
			return nil, err
		}

		encryptedBytes = append(encryptedBytes, encryptedBlockBytes...)
	}

	return encryptedBytes, nil
}

// DecodeOAEP 可变长的RSA数据解密
func (m *RsaEncrypter) DecodeOAEP(ciphertext []byte) ([]byte, error) {
	if m.pri == nil {
		return nil, errors.New("can't decode without private key")
	}
	hash := sha256.New()
	msgLen := len(ciphertext)
	step := m.pub.Size()
	var decryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}
		decryptedBlockBytes, err := rsa.DecryptOAEP(hash, rand.Reader, m.pri, ciphertext[start:finish], nil)
		if err != nil {
			return nil, err
		}

		decryptedBytes = append(decryptedBytes, decryptedBlockBytes...)
	}
	return decryptedBytes, nil
}
