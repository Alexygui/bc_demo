package blockchain

import (
	"github.com/boltdb/bolt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
	"fmt"
)

type Blockchain struct {
	//Blocks []*Block
	Tip []byte //最新区块的hash
	DB  *bolt.DB
}

const dbName = "blockchain.db"
const blockTableName = "blocks"

func CreateGenesisBlockOfBlockchain(data string) {
	if isDBExists() {
		fmt.Println("创始区块已经产生")
		os.Exit(1)
	}

	fmt.Println("正在创建创始区块")
	// 打开或创建数据库
	db, e := bolt.Open(dbName, 0600, nil)
	if e != nil {
		log.Panic(e)
	}

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))

		if nil == bucket {
			bucket, e = tx.CreateBucket([]byte(blockTableName))
			if e != nil {
				log.Panic(e)
			}
		}

		genesisBlock := CreateGenesisBlock(data)
		err := bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = bucket.Put([]byte("tip"), genesisBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

func isDBExists() bool {
	if _, e := os.Stat(dbName); e != nil {
		return false
	}
	return true
}

func ReadGenesisBlock() {
	db, e := bolt.Open(dbName, 0600, nil)
	if e != nil {
		log.Panic(e)
	}
	defer db.Close()

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			tip := bucket.Get([]byte("tip"))
			if tip != nil {
				spew.Dump("b0Hash: ", tip)
				blockData := bucket.Get(tip)
				spew.Dump("blockData: ", blockData)

				block := DeserializeBlock(blockData)
				spew.Dump("block: ", block)
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (bc *Blockchain) AddBlockToBlockchain(data string) {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			blockData := b.Get(bc.Tip)
			preBlock := DeserializeBlock(blockData)

			newBlock := NewBlock(data, preBlock.Height+1, preBlock.Hash)
			e := b.Put(newBlock.Hash, newBlock.Serialize())
			if e != nil {
				log.Panic(e)
			}
			e = b.Put([]byte("tip"), newBlock.Hash)
			bc.Tip = newBlock.Hash
			if e != nil {
				log.Panic(e)
			}

		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (bc *Blockchain) PrintBlockchain() {
	var block *Block

	var blockchainIterator = bc.Iterator()

	for ; blockchainIterator.HasNext(); {
		block = blockchainIterator.Next()
		spew.Dump(block)
	}
}

//产生创始区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 0, []byte{0})
}
