package utils

import "encoding/binary"

func UintToByteArray(val uint64) (buf []byte) {
	buf = make([]byte, 8)
	binary.PutUvarint(buf, val)
	return
}

func ByteArrayToUint64(bytes []byte)  (decode uint64) {
	decode, _ = binary.Uvarint(bytes)
	return
}
