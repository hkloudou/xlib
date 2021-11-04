package xcolor

import (
	"testing"
	//
	// "github.com/stretchr/testify/assert"
)

func TestYellow(t *testing.T) {
	out := Yellow(string([]byte{0x30})) // 0
	t.Log(out)
	// assert.Equal(t, []byte{0x1b, 0x5b, 0x33, 0x33, 0x6d, 0x30, 0x1b, 0x5b, 0x30, 0x6d}, []byte(out))
}

func TestRed(t *testing.T) {
	out := Red("0")
	t.Log(out)
	// assert.Equal(t, []byte{0x1b, 0x5b, 0x33, 0x31, 0x6d, 0x30, 0x1b, 0x5b, 0x30, 0x6d}, []byte(out))
}

func TestBlue(t *testing.T) {
	out := Blue("0")
	t.Log(out)
	// assert.Equal(t, []byte{0x1b, 0x5b, 0x33, 0x34, 0x6d, 0x30, 0x1b, 0x5b, 0x30, 0x6d}, []byte(out))
}

func TestGreen(t *testing.T) {
	out := Green("0")
	t.Log(out)
	// assert.Equal(t, []byte{0x1b, 0x5b, 0x33, 0x32, 0x6d, 0x30, 0x1b, 0x5b, 0x30, 0x6d}, []byte(out))
}
