package vm_test

import (
	"testing"

	"github.com/junbeomlee/vm"
	"github.com/stretchr/testify/assert"
)

func TestSignAndVerify(t *testing.T) {

	pri, pub := vm.GenerateKeyPair()

	digest := []byte("")
	sig := vm.Sign(pri, digest)

	assert.True(t, vm.Verify(pub, sig, digest))
}
