package vm

import "encoding/hex"

func Run(lockingScript string, unlockingScript string, txHash []byte) {

	stack := NewStack()

	ls := parseScript(lockingScript)
	us := parseScript(unlockingScript)

	script := append(ls, us...)

	for _, item := range script {

	}
}

// Parse string script to data and opcode
func parseScript(script string) [][]uint8 {

	codes := make([]Hexable, 0)

	//convert string to hexbytes
	hexBytes, _ := hex.DecodeString(script)

	for i := 0; i < len(hexBytes); i++ {

		hexN := uint8(hexBytes[i])

		//check whether opcode or data
		if hexN > 0 && hexN < OP_PUSHDATA {
			//it is a data
			data := hexBytes[i+1 : i+1+int(hexN)]
			codes = append(codes, data)
			i = i + int(hexN)
			continue
		}

		//it is an opcode
		codes = append(codes, []uint8{GetOpCode(hexN)})
	}

	return codes
}
