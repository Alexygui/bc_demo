package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
	"crypto/sha256"
)

type Block struct {
	//区块高度
	Height int64
	//前一个区块的hash
	PrevBlockHash []byte
	Timestamp     int64
	//交易数据
	Txs    []*Transaction
	Hash   []byte
	Nounce int64
}

func NewBlock(txs []*Transaction, height int64, prevBlockhash []byte) *Block {
	newBlock := &Block{height, prevBlockhash, time.Now().Unix(), txs, nil, 0}
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
//	b.Tip = hash[:]
//}

//序列化block
func (block *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	e := encoder.Encode(block)
	if nil != e {
		panic(e)
	}

	return result.Bytes()
}

//反序列化Block
func DeserializeBlock(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	e := decoder.Decode(&block)
	if nil != e {
		log.Panic(e)
	}
	return &block
}

func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}
