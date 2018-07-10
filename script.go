package vm

import (
	"encoding/hex"
)

// this is a script struct
type Script struct {
	OPcodes [][]uint8
}

func NewScript(script string) Script {

	codes := make([][]uint8, 0)

	hexBytes, _ := hex.DecodeString(script)

	for i := 0; i < len(hexBytes); i++ {

		hexN := uint8(hexBytes[i])

		//check opcode or data
		if hexN > 0 && hexN < OP_PUSHDATA {
			data := hexBytes[i+1 : i+1+int(hexN)]
			codes = append(codes, data)
			i = i + int(hexN)
			continue
		}

		codes = append(codes, []uint8{GetOpCode(hexN)})
	}

	return Script{
		OPcodes: codes,
	}
}
