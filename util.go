package vm

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

func GenerateKeyPair() (ecdsa.PrivateKey, []byte) {

	curve := elliptic.P256()
	private, _ := ecdsa.GenerateKey(curve, rand.Reader)

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}

func Sign(priKey ecdsa.PrivateKey, data []byte) []byte {

	r, s, _ := ecdsa.Sign(rand.Reader, &priKey, data)

	return append(r.Bytes(), s.Bytes()...)
}

func Verify(pubKey []byte, signature []byte, digest []byte) bool {

	curve := elliptic.P256()

	r := big.Int{}
	s := big.Int{}
	sigLen := len(signature)
	r.SetBytes(signature[:(sigLen / 2)])
	s.SetBytes(signature[(sigLen / 2):])

	x := big.Int{}
	y := big.Int{}
	keyLen := len(pubKey)
	x.SetBytes(pubKey[:(keyLen / 2)])
	y.SetBytes(pubKey[(keyLen / 2):])

	rawPubKey := ecdsa.PublicKey{curve, &x, &y}

	if ecdsa.Verify(&rawPubKey, digest, &r, &s) == false {
		return false
	}

	return true
}
