package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/davecgh/go-spew/spew"
	"log"
)

func main() {
	db, e := bolt.Open("github.com/Alexygui/bc_demo/db/my.db", 0600, nil)
	if nil != e {
		log.Panic(e)
	}
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("BlockBucket"))
		if bucket == nil {
			//创建BlockBucket表
			bucket, e = tx.CreateBucket([]byte("BlockBucket"))
			if e != nil {
				return fmt.Errorf("create bucket:%s", e)
			}
		}

		//向表中存数据
		if bucket != nil {
			e := bucket.Put([]byte("tx1"), []byte("send 100 rmb to Bob..."))
			if e != nil {
				log.Panic("数据存储失败。。。")
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("BlockBucket"))
		if nil != bucket {
			data := bucket.Get([]byte("tx1"))
			spew.Dump(data)
			data = bucket.Get([]byte("tx2"))
			spew.Dump(data)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
