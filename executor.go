package vm

import (
	"encoding/hex"
	"errors"
	"fmt"
)

const NONE ScriptType = "NONE"
const P2PK ScriptType = "P2PK"
const P2SH ScriptType = "P2SH"

type ScriptType string

func Run(lockingScript string, unlockingScript string, txHash []byte) (Stack, error) {

	scriptType := CheckScriptType(lockingScript)

	if scriptType == NONE {
		return Stack{}, errors.New("invalid script type")
	}

	if scriptType == P2SH {
		fmt.Println("need to do p2sh")
	}

	fmt.Printf("Script Type is [%s] \n", scriptType)

	stack := NewStack()

	ls := ParseScript(lockingScript)
	us := ParseScript(unlockingScript)

	script := append(ls, us...)

	for _, h := range script {

		opCode, ok := h.(Opcode)

		if ok {
			opCode.Handle(&stack)
		}

		stack.Push(h)
	}

	return stack, nil
}

// Parse string script to data and opcode
func ParseScript(script string) []Hexable {

	asm := make([]Hexable, 0)

	//convert string to hexbytes
	hexBytes, _ := hex.DecodeString(script)

	for i := 0; i < len(hexBytes); i++ {

		hexN := uint8(hexBytes[i])

		//check whether opcode or data
		//hex between 0 and 0x4b is not an opcode but a data, The next byte contains the number of bytes to be pushed onto the stack.
		if hexN > 0 && hexN < OP_PUSHDATA {
			//it is a data
			data := hexBytes[i+1 : i+1+int(hexN)]
			asm = append(asm, Data{Body: data})
			i = i + int(hexN)

			continue
		}

		//it is an opcode
		asm = append(asm, GetOpCode(hexN))
	}

	return asm
}

func CheckScriptType(lockingScript string) ScriptType {

	ls := ParseScript(lockingScript)

	if len(ls) == 2 && ls[1].Hex()[0] == OP_CHECK_SIG {
		return P2PK
	}

	return NONE
}
