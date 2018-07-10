package vm

import (
	"bytes"
	"crypto/sha256"
	"errors"

	"fmt"

	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
)

// Constant
const OP_PUSHDATA uint8 = 0x4c
const OP_TRUE uint8 = 0x51

// Stack
const OP_DUP uint8 = 0x76

// Bitwise logic
const OP_EQUAL uint8 = 0x87
const OP_EQUAL_VERIFY uint8 = 0x88

// Crypto
const OP_CHECK_SIG uint8 = 0xac
const OP_HASH_160 uint8 = 0xa9
const OP_CHECKMULTI_SIG uint8 = 0xae

// can convert data to []uint8 type
type Hexable interface {
	Hex() []uint8
}

type Data struct {
	Body []uint8
}

func (d Data) Hex() []uint8 {
	return d.Body
}

// do something with stack
type StackHandler interface {
	Handle(stack *Stack) error
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
	Opcodes[DupOp{}.Hex()[0]] = DupOp{}
	Opcodes[EqualOp{}.Hex()[0]] = EqualOp{}
	Opcodes[EqualVerifyOp{}.Hex()[0]] = EqualVerifyOp{}
	Opcodes[CheckSigOp{}.Hex()[0]] = CheckSigOp{}
	Opcodes[Hash160Op{}.Hex()[0]] = Hash160Op{}
	Opcodes[CheckMultiSigOp{}.Hex()[0]] = CheckMultiSigOp{}
}

// The variables below refer to the bitcoin opcode [https://en.bitcoin.it/wiki/Script]

// Constant

// desc : The next byte contains the number of bytes to be pushed onto the stack.
type PushDataOp struct {
}

func (PushDataOp) Hex() []uint8 {
	return []uint8{OP_PUSHDATA}
}

func (PushDataOp) Handle(stack *Stack) error {
	//do nothing
	return nil
}

type DupOp struct{}

func (DupOp) Hex() []uint8 {
	return []uint8{OP_DUP}
}

// pop first element and push twice
func (DupOp) Handle(stack *Stack) error {

	h, err := stack.Pop()

	if err != nil {
		return err
	}

	stack.Push(h)
	stack.Push(h)

	return nil
}

// Bitwise logic
type EqualOp struct{}

func (o EqualOp) Hex() []uint8 {
	return []uint8{OP_EQUAL}
}

func (o EqualOp) Handle(stack *Stack) error {

	h1, err := stack.Pop()

	if err != nil {
		return err
	}

	h2, err := stack.Pop()

	if err != nil {
		return err
	}

	if !bytes.Equal(h1.Hex(), h2.Hex()) {
		return errors.New("not equal")
	}

	stack.Push(Data{Body: []uint8{OP_TRUE}})

	return nil
}

type EqualVerifyOp struct{}

func (o EqualVerifyOp) Hex() []uint8 {
	return []uint8{OP_EQUAL_VERIFY}
}

func (o EqualVerifyOp) Handle(stack *Stack) error {

	h1, err := stack.Pop()

	if err != nil {
		return err
	}

	h2, err := stack.Pop()

	if err != nil {
		return err
	}

	if !bytes.Equal(h1.Hex(), h2.Hex()) {
		return errors.New("not equal")
	}

	return nil
}

// Crypto
//
type CheckSigOp struct {
}

func (CheckSigOp) Hex() []uint8 {
	return []uint8{OP_CHECK_SIG}
}

func (CheckSigOp) Handle(stack *Stack) error {
	return nil
}

// OP_HASH_160
type Hash160Op struct {
}

func (Hash160Op) Hex() []uint8 {
	return []uint8{OP_HASH_160}
}

func (Hash160Op) Handle(stack *Stack) error {

	h1, err := stack.Pop()

	if err != nil {
		return err
	}

	dst := make([]byte, hex.EncodedLen(len(h1.Hex())))

	hex.Encode(dst, h1.Hex())

	fmt.Println(dst)

	s := sha256.New()
	s.Write(dst)
	bs := s.Sum(nil)

	r := ripemd160.New()
	r.Write(bs)
	hashed := r.Sum(nil)

	fmt.Println(hashed)

	dst2 := make([]uint8, hex.DecodedLen(len(hashed)))
	hex.Decode(dst2, hashed)
	fmt.Println(dst2)

	stack.Push(Data{Body: hashed})

	fmt.Println(hashed)

	return nil
}

//OP_CHECK_MULTI_SIG_256
type CheckMultiSigOp struct {
}

func (CheckMultiSigOp) Hex() []uint8 {
	return []uint8{OP_CHECKMULTI_SIG}
}

func (CheckMultiSigOp) Handle(stack *Stack) error {
	return nil
}

func GetOpCode(u uint8) Opcode {
	return Opcodes[u]
}
