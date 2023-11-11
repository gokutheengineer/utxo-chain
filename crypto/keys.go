package crypto

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
)

const (
	privKeyLen   = 64
	publicKeyLen = 32
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
