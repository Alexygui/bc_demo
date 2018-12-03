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

// 字节数组反转
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
