package crypto

import (
	"crypto/ed25519"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeyGeneration(t *testing.T) {
	privKey, err := NewPrivateKey()
	require.NoError(t, err)
	require.NotEmpty(t, privKey.key)
	require.Equal(t, len(privKey.key), privKeyLen)

	privKey2, err := NewPrivateKey()
	require.NoError(t, err)
	require.NotEmpty(t, privKey.key)
	require.Equal(t, len(privKey.key), privKeyLen)

	pubKey := privKey.GeneratePublicKey()
	require.NotEmpty(t, pubKey.key)
	require.Equal(t, len(pubKey.key), publicKeyLen)

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
