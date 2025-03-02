package main

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
)

var values = []uint64{
	0,
	1,
	0x203,
	0x4050607,
	0x8090a0b0c0d0e0f,
	0x10,
	0x1112,
	0x13141516,
	0x1718191a1b1c1d1e,
	0x1f202122,
	0x232425262728292a,
	0x2b2c2d2e,
	0x2f303132,
	0x333435363738393a,
	0x3b3c3d3e3f404142,
	0x102030405060708,
	0x123456789abcdef,
	0xfedcba9876543210,
	0xffffffffffffffff,
	0xaaaaaaaaaaaaaaaa,
	0x12345678,
	0x7654321,
	0x1020304,
	0x5060708,
	math.Float64bits(math.Pi),
	math.Float64bits(math.E),
}

// TestPutUint32 verifies that the custom putUint32 function correctly
// writes a uint32 value into a byte slice in little-endian order.
// It compares the output with the [encoding/binary.LittleEndian].PutUint32
// implementation to ensure they produce the same results.
func TestPutUint32(t *testing.T) {
	for _, v := range values {
		v32 := uint32(v)
		b1 := make([]byte, 4)
		b2 := make([]byte, 4)

		putUint32(b1, v32)
		binary.LittleEndian.PutUint32(b2, v32)
		if !bytes.Equal(b1, b2) {
			t.Errorf("putUint32=%v, expect %v", b1, b2)
		}
	}
}
