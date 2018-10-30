package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, e := bolt.Open("github.com/Alexygui/bc_demo/db/my.db", 0600, nil)
	if nil != e {
		log.Panic(e)
	}
	defer db.Close()


}
