package blockchain

import (
	"flag"
	"os"
	"log"
	"fmt"
)

type CLI struct {
	BC *Blockchain
}

func (cli *CLI) RUN() {
	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printCahinCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "addBlockCmd", "添加交易数据")
	createBlockchainData := createBlockchainCmd.String("data", "creating genesis block", "产生创始区块")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printCahinCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			fmt.Println("交易数据不能为空")
			printUsage()
			os.Exit(1)
		}
		//fmt.Println(*flagAddBlockData)
		cli.addBlock(*flagAddBlockData)
	}

	if printCahinCmd.Parsed() {
		fmt.Println("\n输出所有区块的数据：")
		cli.printChain()
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainData == "" {
			fmt.Println("创始区块不可为空")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockOfBlockchain(*createBlockchainData)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.BC.AddBlockToBlockchain(data)
}

//打印区块链中的所有区块
func (cli *CLI) printChain() {
	cli.BC.PrintBlockchain()
}

//产生创始区块并持久化
func (cli *CLI) createGenesisBlockOfBlockchain(data string) {
	CreateGenesisBlockOfBlockchain(data)
	//cli.BC = bc
}

//命令行传入的参数少于2个则打印命令行帮助
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Print(`
Usage:
	addblock -data DATA  --交易数据
	printchain  --打印所有区块信息
	createblockchain -data DATA  --创建创始区块
`)
}
