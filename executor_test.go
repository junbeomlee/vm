package vm_test

import (
	"fmt"
	"testing"

	"github.com/junbeomlee/vm"
)

func TestOpToHex(t *testing.T) {

	var opcode uint8
	opcode = vm.OP_SHA256

	fmt.Printf("%h", opcode)
}
