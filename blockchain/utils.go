package blockchain

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

//int转16进制
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	e := binary.Write(buff, binary.BigEndian, num)
	if e != nil {
		log.Panic(e)
	}
	return buff.Bytes()
}

//json转string数组
func JSONtoArray(jsonStr string) []string {
	var strArr []string
	if err := json.Unmarshal([]byte(jsonStr), &strArr); err != nil {
		log.Panic(err)
	}
	return strArr
}
