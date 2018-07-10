package vm_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/junbeomlee/vm"
)

func TestConvertBytesToHex(t *testing.T) {

	example_script := "4104240ac91558e66c0628693cee5f5120d43caf73cad8586f9f56a447cc6b926520d2b3b259874e5d79dfb4b9aff3405a10cbce47ee820e0824dc7004d5bbcea86fac"

	b, _ := hex.DecodeString(example_script)

	for _, byte := range b {
		hexN := uint8(byte)

		fmt.Println(hexN)

		if hexN == 0x41 {
			fmt.Println("equal")
		}
	}
}

func TestGetOpCode(t *testing.T) {
	example_script := "61"

	b, _ := hex.DecodeString(example_script)

	for _, byte := range b {
		hexN := uint8(byte)
		fmt.Println(vm.GetOpCode(hexN))
	}
}
