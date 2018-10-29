package blockchain

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Height        int64
	PrevBlockHash []byte
	Timestamp     int64
	Data          []byte
	Hash          []byte
}

func NewBlock(data string, height int64, prevBlockhash []byte) *Block {
	newBlock := &Block{height, prevBlockhash, time.Now().Unix(), []byte(data), nil}
	pow := NewProofofWork(newBlock)

	pow.Run()

	newBlock.getHash()
	return newBlock
}

func (b *Block) getHash() {
	heightByts := IntToHex(b.Height)
	timeString := strconv.FormatInt(b.Timestamp, 2)
	timeBytes := []byte(timeString)

	blockBytes := bytes.Join([][]byte{heightByts, b.PrevBlockHash, timeBytes, b.Data}, []byte{})

	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}
