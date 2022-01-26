package utils

import (
	"fmt"
	"math/big"
)

// Signature is signature struct.
type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
