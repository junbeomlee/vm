package ecdsa_test

import (
	"fmt"
	"testing"

	ecdsa2 "crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/junbeomlee/vm/ecdsa"
	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {

	//given
	pri, pub := ecdsa.GetRandomPairKey()

	message := "hello world"
	digest := []byte(message)

	sig, err := ecdsa.Sign(pri, digest)
	assert.NoError(t, err)
	fmt.Println(sig)

	//when
	v, err := ecdsa.Verify(pub, sig, digest)

	//then
	assert.NoError(t, err)
	assert.True(t, v)
}

func TestPriToPEM(t *testing.T) {

	//given
	pri, _ := ecdsa.GetRandomPairKey()

	b, err := ecdsa.PriToPEM(pri)
	assert.NoError(t, err)

	fmt.Printf("%s", b)
}

func TestPubToPEM(t *testing.T) {

	//given
	_, pub := ecdsa.GetRandomPairKey()

	b, err := ecdsa.PubToPEM(pub)
	assert.NoError(t, err)

	fmt.Printf("%s", b)
}

func TestKey(t *testing.T) {

	curve := elliptic.P256()
	private, err := ecdsa2.GenerateKey(curve, rand.Reader)

	assert.NoError(t, err)

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	fmt.Printf("%s", pubKey)
}
