package cli

import (
	"github.com/Alexygui/bc_demo/blockchain"
	"os"
	"fmt"
)

func (cli *CLI) addBlock(txs []*blockchain.Transaction) {
	bc := blockchain.GetBlockchain()
	defer bc.DB.Close()

	bc.AddBlockToBlockchain(txs)
}

//打印区块链中的所有区块
func (cli *CLI) printChain() {
	bc := blockchain.GetBlockchain()
	defer bc.DB.Close()

	bc.PrintBlockchain()
}

//产生创始区块并持久化
func (cli *CLI) createGenesisBlockOfBlockchain(address string) {
	blockchain.CreateGenesisBlockOfBlockchain(address)
}

//命令行传入的参数少于2个则打印命令行帮助
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

//发送交易
func (cli CLI) sendTransaction(from []string, to []string, amount []string) {
	if !blockchain.IsDBExists() {
		fmt.Println("数据不存在")
		os.Exit(1)
	}
	bc := blockchain.BlockchainObject()
	defer bc.DB.Close()

	bc.MineNewBlock(from, to, amount)
}

//查询地址余额
func (cli *CLI) getBalance(address string) {
	fmt.Println("地址：", address)
	//获取blockchain对象
	bc := blockchain.BlockchainObject()
	defer bc.DB.Close()

	amount := bc.GetBalance(address)

	fmt.Printf("%s一共有%d个token\n", address, amount)
}
