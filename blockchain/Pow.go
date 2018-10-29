package blockchain

import "math/big"

const targetBit = 16

type ProofofWork struct {
	Block  *Block
	target *big.Int
}

func NewProofofWork(b *Block) *ProofofWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)
	return &ProofofWork{b, target}
}

func (p *ProofofWork) Run() ([]byte, int64) {
	return nil, 0
}
