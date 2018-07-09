package vm_test

import (
	"testing"
	"github.com/junbeomlee/vm"
	"fmt"
)

func TestOpToHex(t *testing.T){

	var opcode uint8
	opcode = vm.OP_SHA256

	fmt.Printf("%h",opcode)
}


