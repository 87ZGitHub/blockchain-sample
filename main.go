package main

import (
	"blockchain-sample/core"
	"fmt"
)

func main() {
	bc := core.NewBlockChain()

	bc.AddBlock("Send 1 BTC to AShe")
	bc.AddBlock("Send 2 more BTC to AGan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev hash:%x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
