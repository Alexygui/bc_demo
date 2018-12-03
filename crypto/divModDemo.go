package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(100)
	y := big.NewInt(7)
	z, m := big.NewInt(0), big.NewInt(0)
	z.DivMod(x, y, m)
	fmt.Println(x, y, z, m)
}
