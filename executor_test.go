package vm_test

import (
	"fmt"
	"testing"

	"github.com/junbeomlee/vm"
)

func TestOpToHex(t *testing.T) {

	var opcode uint8
	opcode = vm.OP_PUSHDATA

	fmt.Printf("%h", opcode)
}

func TestParseScript(t *testing.T) {
	example_script := "4104240ac91558e66c0628693cee5f5120d43caf73cad8586f9f56a447cc6b926520d2b3b259874e5d79dfb4b9aff3405a10cbce47ee820e0824dc7004d5bbcea86fac"

	asm := vm.ParseScript(example_script)

	for _, h := range asm {
		fmt.Println(h.Hex())
	}
}

func TestRun(t *testing.T) {

	locking_script := "4104240ac91558e66c0628693cee5f5120d43caf73cad8586f9f56a447cc6b926520d2b3b259874e5d79dfb4b9aff3405a10cbce47ee820e0824dc7004d5bbcea86fac"
	unlocking_script := "4730440220277c967dda11986e06e508235006b7e83bc27a1cb0ffaa0d97a543e178199b6a022040d4f8f17865e45de9ca7bcfe3ee2228e175cfcb4468b7650f09b534d3f71f4401"

	vm.Run(locking_script, unlocking_script, []byte{})
}
