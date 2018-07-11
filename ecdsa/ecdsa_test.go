package ecdsa_test

import (
	"fmt"
	"testing"

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
