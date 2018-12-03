package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	s := "alexygui"
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(base64Encoded)

	bytes, e := base64.StdEncoding.DecodeString(base64Encoded)
	if e != nil {
		log.Panic(e)
	}
	fmt.Printf("%x\n", bytes)
	fmt.Println(string(bytes))

	hexStr := hex.EncodeToString([]byte{171, 205})
	fmt.Println(hexStr)

	decodeString, _ := hex.DecodeString("abcd")
	fmt.Printf("%x", decodeString)
}
