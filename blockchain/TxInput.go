package blockchain

type TxInput struct {
	//交易hash
	TxHash []byte
	//UTXO在交易的TXOutput中索引顺序
	Sequence int
	//UTXO的所有者
	ScriptSig string
}