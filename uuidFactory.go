package goflake

import (
	"crypto/rand"
	"encoding/hex"
)

// NewNilUUID returns a UUID with all bits set to zero
func NewNilUUID() *UUID {
	var bytes UUID = [16]byte{}
	for i := 0; i < 16; i++ {
		stringByte, _ := hex.DecodeString("00")
		bytes[i] = stringByte[0]
	}
	return &bytes
}

// NewV4UUID returns a generated version 5 UUID
func NewV4UUID() *UUID {
	bytes := make([]byte, 16)
	rand.Read(bytes)

	result := UUID{}
	for i, v := range bytes {
		result[i] = v
	}

	result.SetVersion(4)
	result.SetVariant()
	return &result
}
