package vm_test

import (
	"crypto/sha256"
	"testing"

	"fmt"

	"encoding/hex"

	"github.com/junbeomlee/vm"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ripemd160"
)

func TestSignAndVerify(t *testing.T) {

	pri, pub := vm.GenerateKeyPair()

	digest := []byte("")
	sig := vm.Sign(pri, digest)

	assert.True(t, vm.Verify(pub, sig, digest))
}

func TestParseScript2(t *testing.T) {

	_, pub := vm.GenerateKeyPair()
	dst := make([]uint8, hex.EncodedLen(len(pub)))
	hex.Encode(dst, pub)

	hexBytes := append(dst, vm.OP_CHECK_SIG)

	asm := make([]vm.Hexable, 0)

	for i := 0; i < len(hexBytes); i++ {

		hexN := uint8(hexBytes[i])

		//check whether opcode or data
		//hex between 0 and 0x4b is not an opcode but a data, The next byte contains the number of bytes to be pushed onto the stack.
		if hexN > 0 && hexN < vm.OP_PUSHDATA {
			//it is a data
			data := hexBytes[i+1 : i+1+int(hexN)]
			asm = append(asm, vm.Data{Body: data})
			i = i + int(hexN)

			continue
		}

		//it is an opcode
		asm = append(asm, vm.GetOpCode(hexN))
	}

	fmt.Println(asm)
}

func TestGenerateKeyPair(t *testing.T) {

	_, pub := vm.GenerateKeyPair()

	fmt.Printf("%s\n", pub)

	dst := make([]uint8, hex.EncodedLen(len(pub)))

	hex.Encode(dst, pub)

	fmt.Printf("%s", dst)
}

func TestHash160(t *testing.T) {

	pubkey := "76a9145e4ff47ceb3a51cdf7ddd80afc4acc5a692dac2d88ac"

	hex.DecodeString()

	s := sha256.New()
	s.Write(pubkey)
	bs := s.Sum(nil)

	r := ripemd160.New()
	r.Write(bs)
	hashed := r.Sum(nil)
}
