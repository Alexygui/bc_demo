package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"encoding/hex"
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
func NewCoinbaseTransaction(address string) *Transaction {
	txInput := &TxInput{[]byte{}, -1, "Genesis Data"}
	txOutput := &TxOutput{10, address}
	txCoinbase := &Transaction{[]byte{}, []*TxInput{txInput}, []*TxOutput{txOutput}}
	txCoinbase.HashTransaction()
	return txCoinbase
}

//对交易进行hash
func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	sum256Bytes := sha256.Sum256(result.Bytes())

	tx.TxHash = sum256Bytes[:]
}

func NewSimpleTransaction(from string, to string, amount int, bc *Blockchain) *Transaction {
	//返回from地址对应的所有的未花费的交易输出
	value, utxosMap := bc.FindEnoughUTXOs(from, amount)

	var txInputs []*TxInput
	var txOutputs []*TxOutput

	for txHash, indexArray := range utxosMap {
		hashBytes, _ := hex.DecodeString(txHash)
		for _, txIndex := range indexArray {
			txInput := &TxInput{hashBytes, txIndex, from}
			txInputs = append(txInputs, txInput)
		}
	}

	txOutput := &TxOutput{int64(amount), to}
	txOutputs = append(txOutputs, txOutput)

	change := value - int64(amount)
	if change > 0 {
		txOutput = &TxOutput{value - int64(amount), from}
		txOutputs = append(txOutputs, txOutput)
	}

	tx := &Transaction{[]byte{}, txInputs, txOutputs}
	tx.HashTransaction()

	return tx
}

//判断交易是否是coinbase交易
func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.TxIn[0].TxHash) == 0 && tx.TxIn[0].Sequence == -1
}
