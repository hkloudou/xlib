package xcert

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"time"
)

type tmplOptions struct {
	pkixName    pkix.Name
	keyUsage    x509.KeyUsage
	extKeyUsage []x509.ExtKeyUsage
	notBefore   time.Duration
	expired     time.Duration
	isCa        bool
	hosts       []string
}
type tmplOption func(*tmplOptions)

func PkixName(name pkix.Name) tmplOption {
	return func(o *tmplOptions) {
		o.pkixName = name
	}
}

func Hosts(hosts ...string) tmplOption {
	return func(o *tmplOptions) {
		o.hosts = hosts
	}
}

func Expired(t time.Duration) tmplOption {
	return func(o *tmplOptions) {
		o.expired = t
	}
}

func NotBefore(t time.Duration) tmplOption {
	return func(o *tmplOptions) {
		o.notBefore = t
	}
}

func IsCa(b bool) tmplOption {
	return func(o *tmplOptions) {
		o.isCa = b
	}
}

func KeyUsage(usg x509.KeyUsage) tmplOption {
	return func(o *tmplOptions) {
		o.keyUsage = usg
	}
}

func ExtKeyUsage(usgs ...x509.ExtKeyUsage) tmplOption {
	return func(o *tmplOptions) {
		o.extKeyUsage = usgs
	}
}

// func Add
