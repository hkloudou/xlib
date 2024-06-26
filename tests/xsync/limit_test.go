//go:build test
// +build test

package xsync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimit(t *testing.T) {
	limit := NewLimit(2)
	limit.Borrow()
	assert.True(t, limit.TryBorrow())
	assert.False(t, limit.TryBorrow())
	assert.Nil(t, limit.Return())
	assert.Nil(t, limit.Return())
	assert.Equal(t, ErrLimitReturn, limit.Return())
}
