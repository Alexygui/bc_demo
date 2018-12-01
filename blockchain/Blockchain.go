package blockchain

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
	"strconv"
)

type Blockchain struct {
	//Blocks []*Block
	Tip []byte //最新区块的hash
	DB  *bolt.DB
}

const dbName = "blockchain.db"
const blockTableName = "blocks"

func CreateGenesisBlockOfBlockchain(address string) {
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

		coinbaseTx := NewCoinbaseTransaction(address)

		genesisBlock := CreateGenesisBlock([]*Transaction{coinbaseTx})
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

func (bc *Blockchain) AddBlockToBlockchain(txs []*Transaction) {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			blockData := b.Get(bc.Tip)
			preBlock := DeserializeBlock(blockData)

			newBlock := NewBlock(txs, preBlock.Height+1, preBlock.Hash)
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
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(txs, 0, []byte{0})
}

//获取保存区块数据的数据库
func GetBlockchain() *Blockchain {
	if !isDBExists() {
		fmt.Println("数据不存在")
		os.Exit(1)
	}
	db, e := bolt.Open(dbName, 0600, nil)
	if e != nil {
		log.Panic(e)
	}

	var tip []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket == nil {
			fmt.Println("尚未生成创始区块")
		}

		tip = bucket.Get([]byte("tip"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{tip, db}
}

func (bc *Blockchain) Iterator() *Iterator {
	return &Iterator{bc.Tip, bc.DB}
}

//从数据库中获取Blockchain对象
func BlockchainObject() *Blockchain {
	db, e := bolt.Open(dbName, 0600, nil)
	if e != nil {
		log.Panic(e)
	}

	var tip []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			tip = b.Get([]byte("tip"))
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{tip, db}
}

//发送交易并打包区块
func (bc *Blockchain) MineNewBlock(from []string, to []string, amount []string) {
	//获取已经存在的最新区块
	var oldBlock *Block
	err := bc.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			oldBlockBytes := bucket.Get(bc.Tip)
			oldBlock = DeserializeBlock(oldBlockBytes)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	var txs []*Transaction
	for i := range from {
		amountI, _ := strconv.Atoi(amount[i])
		tx := NewSimpleTransaction(from[i], to[i], amountI)
		txs = append(txs, tx)
	}

	newBlock := NewBlock(txs, oldBlock.Height+1, oldBlock.Hash)
	//将新区块存储到数据库并更新bc中的tip值
	err = bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket != nil {
			bucket.Put(newBlock.Hash, newBlock.Serialize())
			bucket.Put([]byte("tip"), newBlock.Hash)
			bc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(nil)
	}
}

func (bc *Blockchain) GetBalance(address string) int64 {
	utxos := bc.GetUTXOs(address)

	var amount int64
	for _, v := range utxos {
		amount += int64(v.Output.Value)
	}
	return amount
}

//遍历地址中的Txoutput，查找出未使用的TXOutput，添加到数组中返回
func (bc *Blockchain) GetUTXOs(address string) []*UTXO {
	var UTXOs []*UTXO
	//已经使用的交易输出
	spentTXoutputs := make(map[string][]int)

	iterator := bc.Iterator()
	for ; iterator.HasNext(); {
		block := iterator.Next()

		//遍历当前区块中的所有交易
		for _, tx := range block.Txs {
			//对于非coinbase交易，获取已经使用的TXOutput，
			//根据交易哈希，和在每个交易哈希中的序列放到map中
			if tx.IsCoinbaseTransaction() == false {
				for _, in := range tx.TxIn {
					if in.ScriptSig == address {
						key := hex.EncodeToString(in.TxHash)
						spentTXoutputs[key] = append(spentTXoutputs[key], in.Sequence)
					}
				}
			}

			//遍历当前交易中的TXOutput，将不在已使用交易输出列表中的TXOutput放到UTXO数组中
			for index, txOut := range tx.TxOut {
				if txOut.UnlockScriptPubKeyWithAddress(address) {
					if len(spentTXoutputs) != 0 {
						for txHash, indexArr := range spentTXoutputs {
							for _, i := range indexArr {
								if txHash == hex.EncodeToString(tx.TxHash) && i == index {
									//1.如果当前区块中的某个交易哈希值和已经使用过的交易哈希值相等
									//2.如果已经使用过的TXOutput的序列值和当前交易中的交易序列值相等
									//可以说明这个UTXO是已经被使用过的
									continue
								} else {
									utxo := &UTXO{tx.TxHash, index, txOut}
									UTXOs = append(UTXOs, utxo)
								}
							}
						}
					} else {
						utxo := &UTXO{tx.TxHash, index, txOut}
						UTXOs = append(UTXOs, utxo)
					}
				}
			}
		}
	}

	return UTXOs
}
