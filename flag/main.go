package main

import (
	"flag"
	"os"
	"log"
	"fmt"
)

func main() {
	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printCahinCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "addBlockCmd", "添加交易数据")

	switch os.Args[1] {
	case "addBlock":
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
		fmt.Println(*flagAddBlockData)
	}

	if printCahinCmd.Parsed() {
		fmt.Println("数据所有区块的数据")
	}

	//flagString := flag.String("printchain", "", "输出所有区块信息。。。")
	//flagInt  := flag.Int("number",6,"输出一个证书。。。")
	//flag.Parse()
	//
	//fmt.Printf("%s\n", *flagString)
	//fmt.Printf("%d\n", *flagInt)
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
