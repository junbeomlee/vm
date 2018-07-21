package main

import (
	"fmt"

	"github.com/junbeomlee/vm"
)

// p2pk (Pay to Public key)
// LockingScript=<Public Key> OP_CHECKSIG
// UnLockingScript= <Signature>
func main() {

	result, err := vm.Run(lockingScript, unlockingScript, digest)

	if err != nil {
		panic(err.Error())
	}

	v, err := result.Pop()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("this script is vaild ? [%b]", v)
}

func createLockingScript(pub []byte) string {

	return string(append(pub, vm.OP_CHECK_SIG)[:])
}

func createUnLockingScript(sig []byte) string {
	return string(sig[:])
}
