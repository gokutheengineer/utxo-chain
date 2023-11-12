package crypto

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeyGeneration(t *testing.T) {
	privKey, err := NewPrivateKey()
	require.NoError(t, err)
	require.NotEmpty(t, privKey.key)

	privKey2, err := NewPrivateKey()
	require.NoError(t, err)
	require.NotEmpty(t, privKey.key)

	pubKey := privKey.GeneratePublicKey()
	require.NotEmpty(t, pubKey.key)
	require.Equal(t, len(pubKey.Bytes()), publicKeyLen)

	data := []byte{'a', 'b', 'c'}
	signedData, err := privKey.Sign(data)
	require.NoError(t, err)
	require.NotEmpty(t, signedData)
	require.True(t, ed25519.Verify(pubKey.key, data, signedData))

	signedData, err = privKey2.Sign(data)
	require.NoError(t, err)
	require.NotEmpty(t, signedData)
	require.False(t, ed25519.Verify(pubKey.key, data, signedData))
}

func TestKeyGenerationFromSeed(t *testing.T) {
	seedString := "66a5762137cac3a690e60312a15d2dd99df1e111d9dc6dbb77faa24a7f159e65"
	privKey, err := NewPrivateKeyFromString(seedString)
	require.NoError(t, err)
	require.NotEmpty(t, privKey)

	pubKey := privKey.GeneratePublicKey()
	require.NotEmpty(t, pubKey.key)
	require.Equal(t, len(pubKey.Bytes()), publicKeyLen)

	data := []byte{'a', 'b', 'c'}
	signedData, err := privKey.Sign(data)
	require.NoError(t, err)
	require.NotEmpty(t, signedData)
	require.True(t, ed25519.Verify(pubKey.key, data, signedData))

	seedBytes, err := hex.DecodeString(seedString)
	require.NoError(t, err)
	privKey, err = NewPrivateKeyFromSeed(seedBytes)
	require.NoError(t, err)
	require.NotEmpty(t, privKey)
}
