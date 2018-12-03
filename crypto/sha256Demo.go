package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hasher := sha256.New()
	s := "alexygui"
	hasher.Write([]byte(s))
	bytes := hasher.Sum(nil)
	fmt.Printf("%x",bytes)
}
