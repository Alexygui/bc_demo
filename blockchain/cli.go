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

	flagAddBlockData := addBlockCmd.String("data", "addBlockCmd", "添加交易数据")

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

	default:
		printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
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
}

func (cli *CLI) addBlock(data string)  {
	cli.BC.AddBlockToBlockchain(data)
}

func (cli *CLI) printChain() {
	cli.BC.PrintBlockchain()
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Print(`
Usage:
	addBlock -data DATA -- 交易数据
`)
}
