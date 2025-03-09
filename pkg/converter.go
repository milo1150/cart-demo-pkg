package pkg

import "encoding/binary"

func UintToBytes(n uint) []byte {
	b := make([]byte, 8) // Allocate 8 bytes for uint64
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}
