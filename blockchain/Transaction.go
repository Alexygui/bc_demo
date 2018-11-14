package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//UTXO交易模型
type Transaction struct {
	//交易哈希
	TxHash []byte
	//交易输入
	TxIn []*TxInput
	//交易输出
	TxOut []*TxOutput
}

//创建区块时的coinbase交易
func NewCoinbaseTransaction(address string) *Transaction{
	txInput := &TxInput{[]byte{},-1,"Genesis Data"}
	txOutput := &TxOutput{10,address}
	txCoinbase := &Transaction{[]byte{}, []*TxInput{txInput}, []*TxOutput{txOutput}}
	txCoinbase.HashTransaction()
	return txCoinbase
}

//对交易进行hash
func (tx *Transaction) HashTransaction(){
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if err!= nil {
		log.Panic(err)
	}
	sum256Bytes := sha256.Sum256(result.Bytes())

	tx.TxHash = sum256Bytes[:]
}