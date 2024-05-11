//go:build windows && !appengine
// +build windows,!appengine

package xcolor

import (
	"io"
	"os"
)

// NewColorableStdout returns new instance of writer which handles escape sequence for stdout.
func NewColorableStdout() io.Writer {
	return os.Stdout
}

// NewColorableStderr returns new instance of writer which handles escape sequence for stderr.
func NewColorableStderr() io.Writer {
	return os.Stderr
}
