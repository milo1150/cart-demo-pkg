package pkg

import (
	"encoding/binary"
	"errors"
)

func UintToBytes(n uint) []byte {
	b := make([]byte, 8) // Allocate 8 bytes for uint64
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}

// BytesToUint safely converts a byte slice to a uint value.
func BytesToUint(b []byte) (uint, error) {
	if len(b) == 0 {
		return 0, errors.New("byte slice is empty")
	}

	if len(b) > 8 {
		return 0, errors.New("byte slice exceeds uint64 size")
	}

	// Pad with zeros if the slice is too short
	var padded [8]byte
	copy(padded[8-len(b):], b) // Right-align the bytes (big-endian)

	return uint(binary.BigEndian.Uint64(padded[:])), nil
}
