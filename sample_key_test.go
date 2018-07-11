package vm_test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/junbeomlee/vm"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ripemd160"
)

func TestHash160FromBytes(t *testing.T) {

	b1, _ := hex.DecodeString(vm.PUB_KEY)

	s := sha256.New()
	s.Write(b1)
	bs := s.Sum(nil)

	r := ripemd160.New()
	r.Write(bs)
	hashed := r.Sum(nil)

	fmt.Printf("%x\n", hashed)
	b, _ := hex.DecodeString(vm.HASH_160)

	fmt.Printf("%x", b)
	//assert.Equal(t, hashed, HASH_160)
}

func TestHash160FromUint8s(t *testing.T) {

	hexBytes, _ := hex.DecodeString(vm.PUB_KEY)

	uint8Array := make([]uint8, len(hexBytes))

	for i := 0; i < len(hexBytes); i++ {

		hexN := uint8(hexBytes[i])
		uint8Array[i] = hexN
	}

	s := sha256.New()
	s.Write(uint8Array)
	bs := s.Sum(nil)

	r := ripemd160.New()
	r.Write([]byte(bs))
	hashed := r.Sum(nil)

	fmt.Printf("hash160: %x\n", hashed)
	b, _ := hex.DecodeString(vm.HASH_160)

	fmt.Printf("expected value: %x", b)
	assert.Equal(t, hashed, b)
}
