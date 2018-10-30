package blockchain

import (
	"bytes"
	"fmt"
	"github.com/minio/sha256-simd"
	"math/big"
)

const targetBit = 8

type ProofofWork struct {
	Block  *Block
	target *big.Int
}

func NewProofofWork(b *Block) *ProofofWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)
	return &ProofofWork{b, target}
}

func (pow *ProofofWork) Run() ([]byte, int64) {
	nounce := 0
	var hashInt big.Int
	var hash [32]byte

	for {
		dataBytes := pow.prepareData(nounce)

		hash = sha256.Sum256(dataBytes)
		fmt.Printf("mining: %x\n", hash)
		hashInt.SetBytes(hash[:])

		if pow.target.Cmp(&hashInt) == 1 {
			break
		}

		nounce++
		//time.Sleep(100 * time.Millisecond)
	}

	return hash[:], int64(nounce)
}

func (pow *ProofofWork) prepareData(nounce int) []byte {
	data := bytes.Join(
		[][]byte{
			IntToHex(pow.Block.Height),
			pow.Block.PrevBlockHash,
			IntToHex(pow.Block.Timestamp),
			pow.Block.Data,
			IntToHex(targetBit),
			IntToHex(int64(nounce)),
		},
		[]byte{})
	return data
}
