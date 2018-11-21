package blockchain

type TxOutput struct {
	//交易输出的值
	Value int
	//交易输出的所有者
	ScriptPubKey string
}
//用地址解锁交易输出
func (txOutput *TxOutput) UnlockScriptPubKeyWithAddress(address string) bool {
	return txOutput.ScriptPubKey == address
}