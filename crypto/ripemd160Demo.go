package main

import (
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	hasher := ripemd160.New()
	s := "alexygui"
	hasher.Write([]byte(s))
	bytes := hasher.Sum(nil)
	fmt.Printf("%x", bytes)
}
