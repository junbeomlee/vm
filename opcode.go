package vm

// can convert data to []uint8 type
type Hexable interface {
	Hex() []uint8
}

// do something with stack
type StackHandler interface {
	Handle(stack *Stack)
}

// opcode can do something with stack and convert data to []uint8 type
type Opcode interface {
	Hexable
	StackHandler
}

var Opcodes = make(map[uint8]Opcode, 0)

// init all opcodes
func init() {
	Opcodes[PushDataOp{}.Hex()[0]] = PushDataOp{}
	Opcodes[NopOp{}.Hex()[0]] = NopOp{}
	Opcodes[IfOp{}.Hex()[0]] = IfOp{}
	Opcodes[ElseOp{}.Hex()[0]] = ElseOp{}
	Opcodes[DupOp{}.Hex()[0]] = DupOp{}
	Opcodes[EqualOp{}.Hex()[0]] = EqualOp{}
	Opcodes[EqualVerifyOp{}.Hex()[0]] = EqualVerifyOp{}
	Opcodes[AddOp{}.Hex()[0]] = AddOp{}
	Opcodes[SubOp{}.Hex()[0]] = SubOp{}
	Opcodes[Ripemd160Op{}.Hex()[0]] = Ripemd160Op{}
	Opcodes[CheckSigOp{}.Hex()[0]] = CheckSigOp{}
	Opcodes[Hash160Op{}.Hex()[0]] = Hash160Op{}
	Opcodes[Sha256Op{}.Hex()[0]] = Sha256Op{}
	Opcodes[CheckMultiSigOp{}.Hex()[0]] = CheckMultiSigOp{}
}

// The variables below refer to the bitcoin opcode [https://en.bitcoin.it/wiki/Script]

// Constant

// desc : The next byte contains the number of bytes to be pushed onto the stack.
type PushDataOp struct {
}

func (PushDataOp) Hex() []uint8 {
	return []uint8{0x4b}
}

func (PushDataOp) Handle(stack *Stack) {

}

// Flow Control
// dsec :
type NopOp struct{}

func (NopOp) Hex() []uint8 {
	return []uint8{0x61}
}

func (NopOp) Handle(stack *Stack) {

}

type IfOp struct{}

func (IfOp) Hex() []uint8 {
	return []uint8{0x63}
}

func (IfOp) Handle(stack *Stack) {

}

type ElseOp struct{}

func (ElseOp) Hex() []uint8 {
	return []uint8{0x67}
}

func (ElseOp) Handle(stack *Stack) {}

type DupOp struct{}

func (DupOp) Hex() []uint8 {
	return []uint8{0x76}
}

func (DupOp) Handle(stack *Stack) {}

// Bitwise logic
type EqualOp struct{}

func (o EqualOp) Hex() []uint8 {
	return []uint8{0x87}
}

func (o EqualOp) Handle(stack *Stack) {}

type EqualVerifyOp struct{}

func (o EqualVerifyOp) Hex() []uint8 {
	return []uint8{0x88}
}

func (o EqualVerifyOp) Handle(stack *Stack) {}

// Arithmetic
type AddOp struct{}

func (o AddOp) Hex() []uint8 {
	return []uint8{0x93}
}

func (o AddOp) Handle(stack *Stack) {}

type SubOp struct{}

func (o SubOp) Hex() []uint8 {
	return []uint8{0x94}
}

func (o SubOp) Handle(stack *Stack) {}

// Crypto
type Ripemd160Op struct {
}

func (Ripemd160Op) Hex() []uint8 {
	return []uint8{0xa6}
}

func (Ripemd160Op) Handle(stack *Stack) {}

//
type CheckSigOp struct {
}

func (CheckSigOp) Hex() []uint8 {
	return []uint8{0xac}
}

func (CheckSigOp) Handle(stack *Stack) {}

// OP_HASH_160
type Hash160Op struct {
}

func (Hash160Op) Hex() []uint8 {
	return []uint8{0xa9}
}

func (Hash160Op) Handle(stack *Stack) {}

//OP_SHA256
type Sha256Op struct {
}

func (Sha256Op) Hex() []uint8 {
	return []uint8{0xaa}
}

func (Sha256Op) Handle(stack *Stack) {}

//OP_CHECK_MULTI_SIG_256
type CheckMultiSigOp struct {
}

func (CheckMultiSigOp) Hex() []uint8 {
	return []uint8{0xae}
}

func (CheckMultiSigOp) Handle(stack *Stack) {}

func GetOpCode(u uint8) Opcode {
	return Opcodes[u]
}
