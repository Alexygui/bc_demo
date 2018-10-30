package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Height        int64
	PrevBlockHash []byte
	Timestamp     int64
	Data          []byte
	Hash          []byte
	Nounce        int64
}

func NewBlock(data string, height int64, prevBlockhash []byte) *Block {
	newBlock := &Block{height, prevBlockhash, time.Now().Unix(), []byte(data), nil, 0}
	pow := NewProofofWork(newBlock)

	hash, nounce := pow.Run()
	newBlock.Nounce = nounce
	newBlock.Hash = hash

	//newBlock.getHash()
	return newBlock
}

//func (b *Block) getHash() {
//	heightByts := IntToHex(b.Height)
//	timeString := strconv.FormatInt(b.Timestamp, 2)
//	timeBytes := []byte(timeString)
//
//	blockBytes := bytes.Join([][]byte{heightByts, b.PrevBlockHash, timeBytes, b.Data}, []byte{})
//
//	hash := sha256.Sum256(blockBytes)
//	b.Hash = hash[:]
//}

func (block *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	e := encoder.Encode(block)
	if nil != e {
		panic(e)
	}

	return result.Bytes()
}

func DeserializeBlock(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	e := decoder.Decode(&block)
	if nil != e {
		log.Panic(e)
	}
	return &block
}
