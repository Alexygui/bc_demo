package main

import (
	"github.com/Alexygui/bc_demo/blockchain"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	bc := blockchain.CreateGenesisBlockchain()

	bc.AddBlockToChain("send 1rmb to Alice", bc.Blocks[len(bc.Blocks)-1].Height+1, bc.Blocks[len(bc.Blocks)-1].Hash)

	bc.AddBlockToChain("send 2rmb to Bob", bc.Blocks[len(bc.Blocks)-1].Height+1, bc.Blocks[len(bc.Blocks)-1].Hash)

	bc.AddBlockToChain("send 3rmb to Kat", bc.Blocks[len(bc.Blocks)-1].Height+1, bc.Blocks[len(bc.Blocks)-1].Hash)

	spew.Dump(bc)
	//fmt.Println(bc)
	//fmt.Println(bc.Blocks[0])
}
