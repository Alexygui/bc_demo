package blockchain

import (
	"github.com/boltdb/bolt"
	"log"
	"math/big"
)


type Iterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (blockchainIterator *Iterator) HasNext() bool {
	var hashInt big.Int
	hashInt.SetBytes(blockchainIterator.CurrentHash)

	return big.NewInt(0).Cmp(&hashInt) != 0
}

func (blockchainIterator *Iterator) Next() *Block {
	var block *Block
	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			blockBytes := b.Get(blockchainIterator.CurrentHash)
			block = DeserializeBlock(blockBytes)
			blockchainIterator.CurrentHash = block.PrevBlockHash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}
