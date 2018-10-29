package blockchain

type Blockchain struct {
	Blocks []*Block
}

func CreateGenesisBlockchain() *Blockchain {
	genesisBlock := NewBlock("creating genesis block/生成创始区块...", 0, []byte{0})
	return &Blockchain{[]*Block{genesisBlock}}
}

func (bc *Blockchain) AddBlockToChain(data string, height int64, prevBlockhash []byte) {
	newBlock := NewBlock(data, height, prevBlockhash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
