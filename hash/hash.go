package hash

import (
	"crypto/md5"
	"fmt"
)

// Hash returns the hash value of data.
func Hash(data []byte) uint64 {
	return Sum64(data)
}

// Md5 returns the md5 bytes of data.
func Md5(data []byte) []byte {
	digest := md5.New()
	digest.Write(data)
	return digest.Sum(nil)
}

// Md5Hex returns the md5 hex string of data.
func Md5Hex(data []byte) string {
	return fmt.Sprintf("%x", Md5(data))
}

func Sum64(data []byte) uint64 { return Sum64WithSeed(data, 0) }

// Sum64WithSeed returns the MurmurHash3 sum of data. It is equivalent to the
// following sequence (without the extra burden and the extra allocation):
//
//	hasher := New64WithSeed(seed)
//	hasher.Write(data)
//	return hasher.Sum64()
func Sum64WithSeed(data []byte, seed uint32) uint64 {
	d := &digest128{h1: uint64(seed), h2: uint64(seed)}
	d.seed = seed
	d.tail = d.bmix(data)
	d.clen = len(data)
	h1, _ := d.Sum128()
	return h1
}
