package xcert

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"net"
	"time"
)

/*
生成CA证书
SAMPLE:
func GenarateCA(n pkix.Name) (*Cert, error) {
	if c,err := NewCert(); err != nil {
		return nil, err
	} else if err := c.SignCa(pkix.Name{}); err != nil {
		return nil, err
	} else {
		return c, nil
	}
}
*/
type Cert struct {
	Pub, Pri   bytes.Buffer
	PrivateKey *ecdsa.PrivateKey
	curve      elliptic.Curve
}

func NewCert() (*Cert, error) {
	c := &Cert{}
	if err := c.PrivateKeyGen(); err != nil {
		return nil, err
	}
	return c, nil
}

//PrivateKeyGen 生成PrivateKey
func (m *Cert) PrivateKeyGen() error {
	if m.curve == nil {
		m.curve = elliptic.P256()
	}
	// var err error
	if priKey, err := ecdsa.GenerateKey(m.curve, rand.Reader); err != nil {
		return err
	} else if priKeyEncode, err := x509.MarshalECPrivateKey(priKey); err != nil {
		return err
	} else if err := pem.Encode(&m.Pri, &pem.Block{Type: "EC PRIVATE KEY", Bytes: priKeyEncode}); err != nil {
		return err
	} else {
		m.PrivateKey = priKey
	}
	return nil
}

func (m *Cert) SignCa(n pkix.Name) error {
	if m.PrivateKey == nil {
		return fmt.Errorf("please run keyGen first")
	}
	notBefore := time.Now().Add(-5 * time.Minute).UTC()
	notAfter := notBefore.AddDate(100, 0, 0).UTC()
	if n.CommonName == "" {
		n.CommonName = "root ca cert"
	}
	template := x509.Certificate{
		SerialNumber:          randomSerialNumber(),
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage: x509.KeyUsageDigitalSignature |
			x509.KeyUsageKeyEncipherment | x509.KeyUsageCertSign |
			x509.KeyUsageCRLSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		Subject:     n,
	}
	template.SubjectKeyId = priKeyHash(m.PrivateKey)
	if x509certEncode, err := x509.CreateCertificate(rand.Reader, &template, &template, m.PrivateKey.Public(), m.PrivateKey); err != nil {
		return err
	} else if err := pem.Encode(&m.Pub, &pem.Block{Type: "CERTIFICATE", Bytes: x509certEncode}); err != nil {
		return err
	}
	return nil
}

func (m *Cert) Sign(isServer bool, caCert *x509.Certificate, caKey interface{}, n pkix.Name, dnss []string, ips []net.IP) error {
	if m.PrivateKey == nil {
		return fmt.Errorf("please run keyGen first")
	}

	notBefore := time.Now().Add(-5 * time.Minute).UTC()
	notAfter := notBefore.AddDate(100, 0, 0).UTC()
	//针对服务器模式下的修复
	if isServer {
		if n.CommonName == "" {
			n.CommonName = "localhost"
		} else if len(dnss) == 0 && len(ips) == 0 {
			dnss = []string{"localhost"}
		}
	}
	template := x509.Certificate{
		SerialNumber: randomSerialNumber(),
		NotBefore:    notBefore,
		NotAfter:     notAfter,
		IsCA:         false,
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		Subject:      n,
		IPAddresses:  ips,
		DNSNames:     dnss,
	}
	if isServer {
		template.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	} else {
		template.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}
	}
	template.SubjectKeyId = priKeyHash(m.PrivateKey)
	if x509certEncode, err := x509.CreateCertificate(rand.Reader, &template, caCert, m.PrivateKey.Public(), caKey); err != nil {
		return err
	} else if err := pem.Encode(&m.Pub, &pem.Block{Type: "CERTIFICATE", Bytes: x509certEncode}); err != nil {
		return err
	}
	return nil
}

func ReadPEMCert(pemcert, pemkey []byte) (*x509.Certificate, interface{}, error) {
	cpb, cr := pem.Decode(pemcert)
	fmt.Println(string(cr))
	kpb, kr := pem.Decode(pemkey)
	fmt.Println(string(kr))
	crt, e := x509.ParseCertificate(cpb.Bytes)

	if e != nil {
		fmt.Println("parsex509:", e.Error())
		// os.Exit(1)
		return nil, nil, e
	}
	key, e := x509.ParseECPrivateKey(kpb.Bytes)
	if e != nil {
		fmt.Println("parsekey:", e.Error())
		// os.Exit(1)
		return nil, nil, e
	}
	return crt, key, nil
}
