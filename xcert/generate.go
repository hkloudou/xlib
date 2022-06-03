package xcert

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net"
	"time"
)

func X509Template(opts ...tmplOption) x509.Certificate {
	var options = &tmplOptions{
		isCa:        true,
		expired:     time.Hour * 24 * 365,
		notBefore:   -time.Minute * 5,
		keyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		extKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	for _, o := range opts {
		o(options)
	}
	notBefore := time.Now().Add(options.notBefore)
	notAfter := notBefore.Add(options.expired)
	template := x509.Certificate{
		IsCA:                  options.isCa,
		SerialNumber:          randomSerialNumber(),
		Subject:               options.pkixName,
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              options.keyUsage,
		ExtKeyUsage:           options.extKeyUsage,
		BasicConstraintsValid: true,
	}
	for _, h := range options.hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	return template
}

func X509PemCertGenerate(priv crypto.Signer, tmpl x509.Certificate, usePKCS8 bool, parentCert *x509.Certificate, parentPriv any) ([]byte, []byte, error) {
	// if len(tmpl.SubjectKeyId) == 0 {
	// 	tmpl.SubjectKeyId = priKeyHash(priv)
	// }
	var derBytes []byte
	var err error

	if parentCert != nil && parentPriv != nil {
		derBytes, err = x509.CreateCertificate(rand.Reader, &tmpl, parentCert, priv.Public(), parentPriv)
	} else if parentCert == nil && parentPriv == nil {
		derBytes, err = x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, priv.Public(), priv)
	} else {
		return nil, nil, fmt.Errorf("parentCert and parentPriv should both nil, or both not empty")
	}
	if err != nil {
		return nil, nil, err
	}

	certOut := bytes.NewBuffer(nil)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyOut := bytes.NewBuffer(nil)
	// when force use PKCS8 format
	if usePKCS8 {
		b, err := x509.MarshalPKCS8PrivateKey(priv)
		if err != nil {
			return nil, nil, err
		}
		err = pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: b})
		if err != nil {
			return nil, nil, err
		}
		return certOut.Bytes(), keyOut.Bytes(), nil
	}

	// else
	switch priv := priv.(type) {
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(priv)
		if err != nil {
			return nil, nil, err
		}
		err = pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
		if err != nil {
			return nil, nil, err
		}
		return certOut.Bytes(), keyOut.Bytes(), nil
	case *rsa.PrivateKey:
		b := x509.MarshalPKCS1PrivateKey(priv)
		err = pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: b})
		if err != nil {
			return nil, nil, err
		}
		return certOut.Bytes(), keyOut.Bytes(), nil
	case ed25519.PrivateKey:
		b, err := x509.MarshalPKCS8PrivateKey(priv)
		if err != nil {
			return nil, nil, err
		}
		err = pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: b})
		if err != nil {
			return nil, nil, err
		}
		return certOut.Bytes(), keyOut.Bytes(), nil
	default:
		return nil, nil, errors.New("types are currently supported: *rsa.PrivateKey, *ecdsa.PrivateKey and ed25519.PrivateKey")
	}
}

// func GenerateEcdsaCert(tmpl x509.Certificate) ([]byte, []byte, error) {
// 	// ed25519.GenerateKey()
// 	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	// priv.Public()
// 	tmpl.SubjectKeyId = priKeyHash(priv)
// 	derBytes, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &priv.PublicKey, priv)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	// log.Println("derBytes", len(derBytes))
// 	// create public key
// 	certOut := bytes.NewBuffer(nil)
// 	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

// 	// create private key
// 	keyOut := bytes.NewBuffer(nil)
// 	b, err := x509.MarshalECPrivateKey(priv)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})

// 	return certOut.Bytes(), keyOut.Bytes(), nil
// }

// func generateEcdsaCert(tmpl x509.Certificate) ([]byte, []byte, error) {
// 	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	tmpl.SubjectKeyId = priKeyHash(priv)
// 	derBytes, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &priv.PublicKey, priv)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	// create public key
// 	certOut := bytes.NewBuffer(nil)
// 	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

// 	// create private key
// 	keyOut := bytes.NewBuffer(nil)

// 	b, err := x509.MarshalECPrivateKey(priv)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})

// 	return certOut.Bytes(), keyOut.Bytes(), nil
// }

// func GenerateEcdsaCertWithParent(tmpl x509.Certificate, parentCert *x509.Certificate, parentPriv any) ([]byte, []byte, error) {
// 	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	tmpl.SubjectKeyId = priKeyHash(priv)
// 	derBytes, err := x509.CreateCertificate(rand.Reader, &tmpl, parentCert, &priv.PublicKey, parentPriv)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// create public key
// 	certOut := bytes.NewBuffer(nil)
// 	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

// 	// create private key
// 	keyOut := bytes.NewBuffer(nil)
// 	b, err := x509.MarshalECPrivateKey(priv)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})

// 	return certOut.Bytes(), keyOut.Bytes(), nil
// }

// func EcdsaToPem(cert *x509.Certificate, key *ecdsa.PrivateKey) {
// 	certOut := bytes.NewBuffer(nil)
// 	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

// 	// create private key
// 	keyOut := bytes.NewBuffer(nil)
// 	b, err := x509.MarshalECPrivateKey(priv)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})

// 	return certOut.Bytes(), keyOut.Bytes()
// }
