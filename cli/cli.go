package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/Alexygui/bc_demo/blockchain"
)

type CLI struct {
	BC *blockchain.Blockchain
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

		fromArr := blockchain.JSONtoArray(*flagFrom)
		toArr := blockchain.JSONtoArray(*flagTo)
		amountArr := blockchain.JSONtoArray(*flagAmount)
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
