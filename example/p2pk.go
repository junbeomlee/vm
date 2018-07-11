package main

import (
	"fmt"

	"github.com/junbeomlee/vm"
)

// p2pk (Pay to Public key)
// LockingScript=<Public Key> OP_CHECKSIG
// UnLockingScript= <Signature>
func main() {

	data := []byte("")
	lockingScript := createLockingScript()
	unlockingScript := createUnLockingScript()

	result, err := vm.Run(lockingScript, unlockingScript, data)

	if err != nil {
		panic(err.Error())
	}

	v, err := result.Pop()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("this script is vaild ? [%b]", v)
}

func createLockingScript() string {
	return ""
}

func createUnLockingScript() string {
	return ""
}
