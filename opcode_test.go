package vm_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/junbeomlee/vm"
	"github.com/stretchr/testify/assert"
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

func TestDupOp_Handle(t *testing.T) {

	//given
	dupOp := vm.DupOp{}
	stack := vm.NewStack()
	stack.Push(vm.Data{Body: []uint8{uint8(2)}})

	//when
	dupOp.Handle(&stack)

	//then
	assert.Equal(t, 2, stack.Len())

	h1, err := stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, []uint8{uint8(2)}, h1.Hex())

	h2, err := stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, []uint8{uint8(2)}, h2.Hex())

	assert.Equal(t, 0, stack.Len())
}

func TestEqualOp_Handle(t *testing.T) {

	//given
	equalOp := vm.EqualOp{}
	stack := vm.NewStack()
	stack.Push(vm.Data{Body: []uint8{uint8(2)}})
	stack.Push(vm.Data{Body: []uint8{uint8(2)}})

	//when
	equalOp.Handle(&stack)

	h, err := stack.Pop()
	assert.NoError(t, err)

	assert.Equal(t, h.Hex()[0], vm.OP_TRUE)
}

func TestEqualOp_Handle_ERROR(t *testing.T) {

	//given
	equalOp := vm.EqualOp{}
	stack := vm.NewStack()
	stack.Push(vm.Data{Body: []uint8{uint8(2)}})
	stack.Push(vm.Data{Body: []uint8{uint8(1)}})

	//when
	err := equalOp.Handle(&stack)
	assert.Error(t, err)
}

func TestHash160Op_Handle(t *testing.T) {
	//given
	hashOp := vm.Hash160Op{}
	stack := vm.NewStack()
	stack.Push(vm.Data{Body: []uint8{uint8(2)}})

	//when
	err := hashOp.Handle(&stack)
	assert.NoError(t, err)

	//then
	h, err := stack.Pop()
	assert.NoError(t, err)

	fmt.Printf("%x", h.Hex())
}
