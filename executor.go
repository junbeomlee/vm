package vm

func Run(lockingScript []uint8, unlockingScript []uint8, txHash []byte) {

	script := append(unlockingScript, lockingScript...)

	for _, byte := range script {
		switch byte {
		case OP_PUSHDATA:

		case OP_NOP:
		case OP_IF:
		case OP_ELSE:
		case OP_DUP:
		case
		}
	}
}
