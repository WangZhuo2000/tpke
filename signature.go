package tpke

import "github.com/bls"

type Signature struct {
	G2 *bls.G2Projective
}

func (s *Signature) parity() bool {
	return false
}
