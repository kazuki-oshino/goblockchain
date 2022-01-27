package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
)

// Signature is signature struct.
type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%064x%064x", s.R, s.S)
}

// String2BigIntTuple is convert string to big int.
func String2BigIntTuple(s string) (big.Int, big.Int) {
	bx, _ := hex.DecodeString(s[:64])
	by, _ := hex.DecodeString(s[64:])

	var bix, biy big.Int
	_ = bix.SetBytes(bx)
	_ = biy.SetBytes(by)
	return bix, biy
}

// SignatureFromString is convert string to Signature.
func SignatureFromString(s string) *Signature {
	x, y := String2BigIntTuple(s)
	return &Signature{R: &x, S: &y}
}

// PublicKeyFromString is convert string to PublicKey.
func PublicKeyFromString(s string) *ecdsa.PublicKey {
	x, y := String2BigIntTuple(s)
	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: &x, Y: &y}
}

// PrivateKeyFromString is conver string to PrivateKey.
func PrivateKeyFromString(s string, publicKey *ecdsa.PublicKey) *ecdsa.PrivateKey {
	b, _ := hex.DecodeString(s[:])
	var bi big.Int
	_ = bi.SetBytes(b)
	return &ecdsa.PrivateKey{PublicKey: *publicKey, D: &bi}
}
