package xsys

import (
	"os"

	"github.com/hkloudou/xlib/xstring"
)

var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = xstring.RandId()
	}
}

// Hostname returns the name of the host, if no hostname, a random id is returned.
func Hostname() string {
	return hostname
}
