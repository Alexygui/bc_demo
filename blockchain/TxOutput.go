package blockchain

type TxOutput struct {
	//交易输出的值
	Value int
	//交易输出的所有者
	ScriptPubKey string
}