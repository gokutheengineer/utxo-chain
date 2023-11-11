package crypto

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const (
	privKeyLen   = 64
	publicKeyLen = 32
	addressLen   = 20
	seedSize     = 32
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func (pk *PrivateKey) Bytes() []byte {
	return pk.key
}

func (pk *PrivateKey) Sign(message []byte) (signedData []byte, err error) {
	signedData, err = pk.key.Sign(rand.Reader, message, crypto.Hash(0))
	if err != nil {
		return
	}
	return
}

func NewPrivateKey() (*PrivateKey, error) {
	_, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &PrivateKey{key: privKey}, nil
}

func NewPrivateKeyFromSeed(seed []byte) (*PrivateKey, error) {
	if len(seed) != seedSize {
		return nil, fmt.Errorf("provided seed's size: %d doesn't match the required seed size: %d", len(seed), seedSize)
	}
	privKey := ed25519.NewKeyFromSeed(seed)
	return &PrivateKey{key: privKey}, nil
}

func NewPrivateKeyFromString(seedString string) (*PrivateKey, error) {
	seed, err := hex.DecodeString(seedString)
	if err != nil {
		return nil, err
	}
	return NewPrivateKeyFromSeed(seed)
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (pk *PrivateKey) GeneratePublicKey() *PublicKey {
	pubKey := pk.key.Public().(ed25519.PublicKey)
	return &PublicKey{key: pubKey}
}

func (pk *PublicKey) Bytes() []byte {
	return pk.key
}

func (pk *PublicKey) VerifySignature(msg, signedMsg []byte) bool {
	return ed25519.Verify(pk.key, msg, signedMsg)
}

func (pk *PublicKey) Address() Address {
	return Address{
		value: pk.key[publicKeyLen-addressLen:],
	}
}

type Address struct {
	value []byte
}

func (a Address) String() string {
	return hex.EncodeToString(a.value)
}
