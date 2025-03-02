package main

import "unsafe"

// Like [encoding/binary.LittleEndian].PutUint32, but in one line.
func putUint32(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	*(*uint32)(unsafe.Pointer(&b[0])) = v
}
