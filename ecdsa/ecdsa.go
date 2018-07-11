package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"math/big"
)

type EcdsaSignature struct {
	R, S *big.Int
}

func UnmarshalECDSASignature(signature []byte) (*big.Int, *big.Int, error) {

	ecdsaSig := new(EcdsaSignature)
	_, err := asn1.Unmarshal(signature, ecdsaSig)

	if err != nil {
		return nil, nil, errors.New("failed to unmarshal")
	}

	if ecdsaSig.R == nil {
		return nil, nil, errors.New("invalid signature")
	}
	if ecdsaSig.S == nil {
		return nil, nil, errors.New("invalid signature")
	}

	if ecdsaSig.R.Sign() != 1 {
		return nil, nil, errors.New("invalid signature")
	}
	if ecdsaSig.S.Sign() != 1 {
		return nil, nil, errors.New("invalid signature")
	}

	return ecdsaSig.R, ecdsaSig.S, nil
}

func Decode(pemEncodedPub []uint8) (*ecdsa.PublicKey, error) {

	blockPub, _ := pem.Decode(pemEncodedPub)

	x509EncodedPub := blockPub.Bytes
	genericPublicKey, err := x509.ParsePKIXPublicKey(x509EncodedPub)

	if err != nil {
		return nil, err
	}

	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return publicKey, nil
}

func Encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, []byte) {

	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return pemEncoded, pemEncodedPub
}

func Sign(privateKey *ecdsa.PrivateKey, digest []byte) ([]byte, error) {

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, digest)

	if err != nil {
		return nil, err
	}

	signature, err := marshalECDSASignature(r, s)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func marshalECDSASignature(r, s *big.Int) ([]byte, error) {
	return asn1.Marshal(EcdsaSignature{r, s})
}

// Sign signs a digest(hash) using priKey(private key), and returns signature.
func Verify(pubKey *ecdsa.PublicKey, signature, digest []byte) (bool, error) {

	r, s, err := UnmarshalECDSASignature(signature)

	if err != nil {
		return false, err
	}

	valid := ecdsa.Verify(pubKey, digest, r, s)

	if !valid {
		return valid, errors.New("failed to verify")
	}

	return valid, nil

}

func GetRandomPairKey() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey
}

func PriToPEM(key *ecdsa.PrivateKey) ([]byte, error) {

	keyData, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "ECDSA PRIVATE KEY",
			Bytes: keyData,
		},
	), nil
}

func PubToPEM(key *ecdsa.PublicKey) ([]byte, error) {
	keyData, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "ECDSA PUBLIC KEY",
			Bytes: keyData,
		},
	), nil
}
