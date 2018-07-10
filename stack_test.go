package vm_test

import (
	"testing"

	"github.com/junbeomlee/vm"
	"github.com/stretchr/testify/assert"
)

func TestStack_PushAndPop(t *testing.T) {

	stack := vm.NewStack()

	for i := 0; i < 10; i++ {
		stack.Push(vm.Data{Body: []uint8{uint8(i)}})
	}

	for i := 0; i < 10; i++ {
		v, err := stack.Pop()
		assert.NoError(t, err)
		assert.Equal(t, v.Hex(), []uint8{uint8(9 - i)})
	}
}
