package xcert

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

func ParseCertPair(certParame, keyParame interface{}) (*x509.Certificate, crypto.PrivateKey, error) {
	fail := func(err error) (*x509.Certificate, crypto.PrivateKey, error) { return nil, nil, err }
	_, cert, err := ParsePublicCert(certParame)
	if err != nil {
		return fail(err)
	}
	key, err := ParsePrivateCert(keyParame)
	if err != nil {
		return fail(err)
	}
	return cert, key, nil
}

/*
Public Key
*/
//ParsePublicCert []byte string *x509.Certificate
func ParsePublicCert(cert interface{}) ([][]byte, *x509.Certificate, error) {
	switch cert.(type) {
	case string:
		if cert.(string) == "" {
			return nil, nil, nil
		}
		bt, err := ioutil.ReadFile(cert.(string))
		if err != nil {
			return nil, nil, err
		}
		return readPemPublicCert(bt)
	case []byte:
		if len(cert.([]byte)) == 0 {
			return nil, nil, nil
		}
		return readPemPublicCert(cert.([]byte))
	// case *x509.Certificate:
	// 	return nil, (cert.(*x509.Certificate)), nil
	case nil:
		return nil, nil, nil
	default:
		return nil, nil, fmt.Errorf("xcert: failed to ReadPublicCert cert type:%v", reflect.TypeOf(cert))
	}
}

//Copy From package tls
func readPemPublicCert(certPEMBlock []byte) ([][]byte, *x509.Certificate, error) {
	fail := func(err error) ([][]byte, *x509.Certificate, error) { return nil, nil, err }
	var skippedBlockTypes []string
	var certs = make([][]byte, 0)
	for {
		var certDERBlock *pem.Block
		certDERBlock, certPEMBlock = pem.Decode(certPEMBlock)
		if certDERBlock == nil {
			break
		}
		if certDERBlock.Type == "CERTIFICATE" {
			certs = append(certs, certDERBlock.Bytes)
		} else {
			skippedBlockTypes = append(skippedBlockTypes, certDERBlock.Type)
		}
	}

	if len(certs) == 0 {
		if len(skippedBlockTypes) == 0 {
			return fail(errors.New("xcert: failed to find any PEM data in certificate input"))
		}
		if len(skippedBlockTypes) == 1 && strings.HasSuffix(skippedBlockTypes[0], "PRIVATE KEY") {
			return fail(errors.New("xcert: failed to find certificate PEM data in certificate input, but did find a private key; PEM inputs may have been switched"))
		}
		return fail(fmt.Errorf("xcert: failed to find \"CERTIFICATE\" PEM block in certificate input after skipping PEM blocks of the following types: %v", skippedBlockTypes))
	}
	tmp, err := x509.ParseCertificate(certs[0])
	return certs, tmp, err
}

/*
Private Key
*/
//ReadPrivateCert []byte string *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey
func ParsePrivateCert(cert interface{}) (crypto.PrivateKey, error) {
	switch cert.(type) {
	case string:
		bt, err := ioutil.ReadFile(cert.(string))
		if err != nil {
			return nil, err
		}
		return readPemPrivateCert(bt)
	case []byte:
		return readPemPrivateCert(cert.([]byte))
	case *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey:
		return cert, nil
	default:
		return nil, fmt.Errorf("xcert: failed to ReadPrivateCert cert type:%v", reflect.TypeOf(cert))
	}
}
func readPemPrivateCert(keyPEMBlock []byte) (crypto.PrivateKey, error) {
	fail := func(err error) (interface{}, error) { return nil, err }
	var skippedBlockTypes []string
	var keyDERBlock *pem.Block
	for {
		keyDERBlock, keyPEMBlock = pem.Decode(keyPEMBlock)
		if keyDERBlock == nil {
			if len(skippedBlockTypes) == 0 {
				return fail(errors.New("xcert: failed to find any PEM data in key input"))
			}
			if len(skippedBlockTypes) == 1 && skippedBlockTypes[0] == "CERTIFICATE" {
				return fail(errors.New("xcert: found a certificate rather than a key in the PEM for the private key"))
			}
			return fail(fmt.Errorf("xcert: failed to find PEM block with type ending in \"PRIVATE KEY\" in key input after skipping PEM blocks of the following types: %v", skippedBlockTypes))
		}
		if keyDERBlock.Type == "PRIVATE KEY" || strings.HasSuffix(keyDERBlock.Type, " PRIVATE KEY") {
			break
		}
		skippedBlockTypes = append(skippedBlockTypes, keyDERBlock.Type)
	}
	privateKey, err := parsePrivateKey(keyDERBlock.Bytes)
	if err != nil {
		return fail(err)
	}
	return privateKey, nil
}

// Copy From package tls
// Attempt to parse the given private key DER block. OpenSSL 0.9.8 generates
// PKCS #1 private keys by default, while OpenSSL 1.0.0 generates PKCS #8 keys.
// OpenSSL ecparam generates SEC1 EC private keys for ECDSA. We try all three.
func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey:
			return key, nil
		default:
			return nil, errors.New("xcert: found unknown private key type in PKCS#8 wrapping")
		}
	}
	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}

	return nil, errors.New("xcert: failed to parse private key")
}

func ParseTlsConfig(caParame interface{}, certParame interface{}, keyParame interface{}) (*tls.Config, error) {
	fail := func(err error) (*tls.Config, error) { return nil, err }
	var pool *x509.CertPool
	cfg := &tls.Config{}
	cfg.MinVersion = tls.VersionTLS13
	_, ca, err := ParsePublicCert(caParame)
	if err != nil {
		return fail(err)
	}
	if ca != nil {
		if !ca.IsCA {
			return fail(fmt.Errorf("xcert: ca should have IsCa=true"))
		}
		pool = x509.NewCertPool()
		pool.AddCert(ca)
	}
	certBytes, cert, err := ParsePublicCert(certParame)
	if err != nil {
		return fail(err)
	}
	if cert == nil {
		return fail(fmt.Errorf("xcert: cert parse error"))
	}
	if pool == nil {
		if !cert.IsCA {
			return fail(fmt.Errorf("xcert: Cert should have IsCa=true when not contain ca"))
		}
	} else {
		if cert.IsCA {
			return fail(fmt.Errorf("xcert: Cert should have IsCa=false when contain ca"))
		}
	}

	//if have ca pool
	if pool != nil {
		for i := 0; i < len(cert.ExtKeyUsage); i++ {
			usage := cert.ExtKeyUsage[i]
			// log.Println("ext", usage)
			switch usage {
			case x509.ExtKeyUsageClientAuth:
				cfg.RootCAs = pool
				break
			case x509.ExtKeyUsageServerAuth:
				cfg.ClientCAs = pool
				cfg.ClientAuth = tls.RequireAndVerifyClientCert
				break
			case x509.ExtKeyUsageAny:
				cfg.RootCAs = pool
				cfg.ClientCAs = pool
				break
			}
		}
		if cfg.ClientCAs == nil && cfg.RootCAs == nil {
			return fail(fmt.Errorf("xcert: Cert ExtKeyUsage neither contain ExtKeyUsageServerAuth nor ExtKeyUsageClientAuth"))
		}
	} else {
		//skip verify if not have ca cert
		cfg.InsecureSkipVerify = true
	}

	key, err := ParsePrivateCert(keyParame)
	if err != nil {
		return fail(err)
	}
	cfg.Certificates = append(cfg.Certificates, tls.Certificate{
		Certificate: certBytes,
		PrivateKey:  key,
	})
	return cfg, nil
}
