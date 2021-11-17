package main

import (
	"archive/zip"
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		// log.Println(err)
		return false
	}
	return s.IsDir()
}

func readPEMCert(pemcert, pemkey []byte) (*x509.Certificate, interface{}, error) {
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

func packZip(name string, ca, cert, key []byte) ([]byte, error) {
	var zipbuf bytes.Buffer
	zipWriter := zip.NewWriter(&zipbuf)
	if a, err := zipWriter.Create("ca.pem"); err != nil {
		return nil, err
	} else if _, err := a.Write(ca); err != nil {
		return nil, err
	}

	if a, err := zipWriter.Create(name + ".pem"); err != nil {
		return nil, err
	} else if _, err := a.Write(cert); err != nil {
		return nil, err
	}

	if a, err := zipWriter.Create(name + ".key"); err != nil {
		return nil, err
	} else if _, err := a.Write(key); err != nil {
		return nil, err
	}
	zipWriter.Close()
	return zipbuf.Bytes(), nil
}
