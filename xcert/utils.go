package xcert

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"net"
	"time"
)

// func ReadEcdsaPemCert(pemcert, pemkey []byte) (*x509.Certificate, *ecdsa.PrivateKey, error) {
// 	cpb, cr := pem.Decode(pemcert)
// 	fmt.Println(string(cr))
// 	kpb, kr := pem.Decode(pemkey)
// 	fmt.Println(string(kr))
// 	crt, e := x509.ParseCertificate(cpb.Bytes)

// 	if e != nil {
// 		return nil, nil, e
// 	}
// 	key, e := x509.ParseECPrivateKey(kpb.Bytes)
// 	if e != nil {
// 		return nil, nil, e
// 	}
// 	return crt, key, nil
// }

func Template(opts ...tmplOption) x509.Certificate {
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

func GenerateEcdsaCert(tmpl x509.Certificate) ([]byte, []byte, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	tmpl.SubjectKeyId = priKeyHash(priv)
	derBytes, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}

	// create public key
	certOut := bytes.NewBuffer(nil)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	// create private key
	keyOut := bytes.NewBuffer(nil)
	b, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})

	return certOut.Bytes(), keyOut.Bytes(), nil
}

func GenerateEcdsaCertWithParent(tmpl x509.Certificate, parentCert *x509.Certificate, parentPriv any) ([]byte, []byte, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	tmpl.SubjectKeyId = priKeyHash(priv)
	derBytes, err := x509.CreateCertificate(rand.Reader, &tmpl, parentCert, &priv.PublicKey, parentPriv)
	if err != nil {
		return nil, nil, err
	}

	// create public key
	certOut := bytes.NewBuffer(nil)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	// create private key
	keyOut := bytes.NewBuffer(nil)
	b, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})

	return certOut.Bytes(), keyOut.Bytes(), nil
}
