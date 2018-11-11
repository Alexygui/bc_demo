package main

import "github.com/Alexygui/bc_demo/blockchain"

func main() {
	bc := blockchain.CreateGenesisBlockBoltDB()
	defer bc.DB.Close()
	////blockchain.ReadGenesisBlock()

	bc.AddBlockToBlockchain("send 100RMB to Alice")
	bc.AddBlockToBlockchain("send 200RMB to Bob")
	bc.AddBlockToBlockchain("send 300RMB to Kat")

	bc.PrintBlockchain()
}

//func seriable_deseriable(){
//
//	bc := blockchain.CreateGenesisBlockchain()
//
//	bc.AddBlockToChain("send 1rmb to Alice", bc.Blocks[len(bc.Blocks)-1].Height+1, bc.Blocks[len(bc.Blocks)-1].Tip)
//
//	bc.AddBlockToChain("send 2rmb to Bob", bc.Blocks[len(bc.Blocks)-1].Height+1, bc.Blocks[len(bc.Blocks)-1].Tip)
//
//	bc.AddBlockToChain("send 3rmb to Kat", bc.Blocks[len(bc.Blocks)-1].Height+1, bc.Blocks[len(bc.Blocks)-1].Tip)
//
//	fmt.Println(bc)
//	fmt.Println(bc.Blocks[0])
//
//	block := bc.Blocks[0]
//	spew.Dump(block)
//
//	bytes := block.Serialize()
//	spew.Dump(bytes)
//
//	block = blockchain.DeserializeBlock(bytes)
//	spew.Dump(block)
//}
//
////序列化，存储如boltdb，取出，反序列化
//func db_seriable_deseriable(){
//	bc := blockchain.CreateGenesisBlockchain()
//
//	//向bolt数据库中存入序列化后的数据，读取数据进行反序列化成block结构体
//	block := bc.Blocks[0]
//	spew.Dump(block)
//	bytes := block.Serialize()
//	spew.Dump(bytes)
//	db, err := bolt.Open("github.com/Alexygui/bc_demo/db/my.db", 0600, nil)
//	if err != nil {
//		log.Panic(err)
//	}
//	defer db.Close()
//	err = db.Update(func(tx *bolt.Tx) error {
//		bucket := tx.Bucket([]byte("blocks"))
//		if bucket == nil {
//			createBucket, e := tx.CreateBucket([]byte("blocks"))
//			if e != nil {
//				log.Panic(e)
//			}
//			err = createBucket.Put([]byte("b0"), block.Serialize())
//			if err != nil {
//				log.Panic(err)
//			}
//		}
//		return nil
//	})
//	if err != nil {
//		log.Panic(err)
//	}
//
//	err = db.View(func(tx *bolt.Tx) error {
//		bucket := tx.Bucket([]byte("blocks"))
//		if bucket != nil {
//			data := bucket.Get([]byte("b0"))
//			spew.Dump(data)
//			block = blockchain.DeserializeBlock(data)
//			spew.Dump(block)
//		}
//		return nil
//	})
//	if err != nil {
//		log.Panic(err)
//	}
//}
