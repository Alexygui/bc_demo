package blockchain

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	BC *Blockchain
}

//打印命令行帮助
func printUsage() {
	fmt.Print(`
Usage:
	sendTransaction -from FROM -to TO -amount AMOUNT  --交易数据
	printchain  --打印所有区块信息
	createBlockchain -address ADDRESS  --创建创始区块
	getbalance -address ADDRESS  --获取账户余额
`)
}

func (cli *CLI) RUN() {
	isValidArgs()

	// e.g.:  sendTransaction -from '["Alice","Alex"]' -to '["Bob","Bag"]' -amount '["10","20"]'
	sendTransactionCmd := flag.NewFlagSet("sendTransaction", flag.ExitOnError)
	printCahinCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createBlockchain", flag.ExitOnError)
	getbalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	flagFrom := sendTransactionCmd.String("from", "", "转账源地址")
	flagTo := sendTransactionCmd.String("to", "", "转账目标地址")
	flagAmount := sendTransactionCmd.String("amount", "", "添加交易数据")

	flagCreateBlockchainWithAddress := createBlockchainCmd.String("address", "", "设置产生创始区块的地址")
	flagGetbalanceWithAddress := getbalanceCmd.String("address", "", "获取某个地址的余额")

	switch os.Args[1] {
	case "sendTransaction":
		err := sendTransactionCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printCahinCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createBlockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getbalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		printUsage()
		os.Exit(1)
	}

	if sendTransactionCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			fmt.Println("交易数据不能为空")
			printUsage()
			os.Exit(1)
		}
		//fmt.Println(*flagFrom, *flagTo, *flagAmount)
		//fmt.Println(JSONtoArray(*flagFrom),JSONtoArray(*flagTo),JSONtoArray(*flagAmount))

		fromArr := JSONtoArray(*flagFrom)
		toArr := JSONtoArray(*flagTo)
		amountArr := JSONtoArray(*flagAmount)
		cli.sendTransaction(fromArr, toArr, amountArr)
	}

	if printCahinCmd.Parsed() {
		fmt.Println("\n输出所有区块的数据：")
		cli.printChain()
	}

	if createBlockchainCmd.Parsed() {
		if *flagCreateBlockchainWithAddress == "" {
			fmt.Println("创始区块不可为空")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockOfBlockchain(*flagCreateBlockchainWithAddress)
	}

	if getbalanceCmd.Parsed() {
		if *flagGetbalanceWithAddress == "" {
			fmt.Println("地址不可为空")
			printUsage()
			os.Exit(1)
		}
		cli.getBalance(*flagGetbalanceWithAddress)
	}
}

func (cli *CLI) addBlock(txs []*Transaction) {
	bc := GetBlockchain()
	defer bc.DB.Close()

	bc.AddBlockToBlockchain(txs)
}

//打印区块链中的所有区块
func (cli *CLI) printChain() {
	bc := GetBlockchain()
	defer bc.DB.Close()

	bc.PrintBlockchain()
}

//产生创始区块并持久化
func (cli *CLI) createGenesisBlockOfBlockchain(address string) {
	CreateGenesisBlockOfBlockchain(address)
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
	if !isDBExists() {
		fmt.Println("数据不存在")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()

	blockchain.MineNewBlock(from, to, amount)
}

//查询地址余额
func (cli *CLI) getBalance(address string) {
	fmt.Println("地址：", address)
	unspentTXs := UnspentTransactionsWithAddress(address)
	fmt.Println(unspentTXs)
}
