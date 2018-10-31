package blockchain

import (
	"github.com/boltdb/bolt"
	"github.com/davecgh/go-spew/spew"
	"log"
)

type Blockchain struct {
	//Blocks []*Block
	Hash []byte
	DB   *bolt.DB
}

const dbName = "github.com/Alexygui/bc_demo/db/blockchain.db"
const blockTableName = "blocks"

//func CreateGenesisBlockchain() *Blockchain {
//	genesisBlock := NewBlock("creating genesis block/生成创始区块...", 0, []byte{0})
//	return &Blockchain{[]*Block{genesisBlock}}
//}
//
//func (bc *Blockchain) AddBlockToChain(data string, height int64, prevBlockhash []byte) {
//	newBlock := NewBlock(data, height, prevBlockhash)
//	bc.Blocks = append(bc.Blocks, newBlock)
//}

func CreateGenesisBlockBoltDB() *Blockchain {
	db, e := bolt.Open(dbName, 0600, nil)
	if e != nil {
		log.Panic(e)
	}

	var blockHash []byte

	err := db.Update(func(tx *bolt.Tx) error {
		bucket, e := tx.CreateBucket([]byte(blockTableName))
		if e != nil {
			log.Panic(e)
		}

		if bucket != nil {
			genesisBlock := NewBlock("creating genesis block/生成创始区块...", 0, []byte{0})
			err := bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = bucket.Put([]byte("b0"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			blockHash = genesisBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{blockHash, db}
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
			b0 := bucket.Get([]byte("b0"))
			if b0 != nil {
				spew.Dump("b0Hash: ", b0)
				blockData := bucket.Get(b0)
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
			blockData := b.Get(bc.Hash)
			preBlock := DeserializeBlock(blockData)

			newBlock := NewBlock(data, preBlock.Height+1, preBlock.Hash)
			e := b.Put(newBlock.Hash, newBlock.Serialize())
			if e != nil {
				log.Panic(e)
			}
			e = b.Put([]byte("b0"), newBlock.Hash)
			if e!= nil{
				log.Panic(e)
			}

		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
