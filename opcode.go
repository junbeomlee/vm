package vm

// The variables below refer to the bitcoin opcode [https://en.bitcoin.it/wiki/Script]

// Constant
var OP_PUSHDATA uint8 = 0x4c

// Flow Control
var OP_NOP uint8 = 0x61
var OP_IF uint8 = 0x63
var OP_ELSE uint8 = 0x67

// Stack
var OP_DUP uint8 = 0x76

// Bitwise logic
var OP_EQUAL uint8 = 0x87
var OP_EQUAL_VERIFY uint8 = 0x88

// Arithmetic
var OP_ADD uint8 = 0x93
var OP_SUB uint8 = 0x94

// Crypto
var OP_RIPEMD160 uint8 = 0xa6
var OP_CHECK_SIG uint8 = 0xac
var OP_HASH_160 uint8 = 0xa8
var OP_SHA256 uint8 = 0xaa
