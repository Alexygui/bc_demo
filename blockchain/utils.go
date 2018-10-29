package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	e := binary.Write(buff, binary.BigEndian, num)
	if e != nil {
		log.Panic(e)
	}
	return buff.Bytes()
}
