package vm_test

import (
	"testing"

	"fmt"

	"github.com/junbeomlee/vm"
	"github.com/stretchr/testify/assert"
)

func TestNewScript(t *testing.T) {

	// this script code is [pubkey + checkSig] "4104240ac91558e66c0628693cee5f5120d43caf73cad8586f9f56a447cc6b926520d2b3b259874e5d79dfb4b9aff3405a10cbce47ee820e0824dc7004d5bbcea86fac"
	example_script := "4104240ac91558e66c0628693cee5f5120d43caf73cad8586f9f56a447cc6b926520d2b3b259874e5d79dfb4b9aff3405a10cbce47ee820e0824dc7004d5bbcea86fac"

	script := vm.NewScript(example_script)

	assert.Equal(t, 0xac, int(script.OPcodes[1][0]))
	fmt.Println(script.OPcodes)
}
