package util

import "golang.org/x/crypto/blake2b"

// HashBlake2b hashes the input with BLAKE2b-256
func HashBlake2b(input []byte) []byte {
	hash := blake2b.Sum256(input)
	return hash[:]
}
