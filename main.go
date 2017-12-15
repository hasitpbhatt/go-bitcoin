package main

import (
	"fmt"

	"github.com/hasitpbhatt/go-bitcoin/blockchain"
)

func main() {
	chain := blockchain.NewBlockChain()
	chain.AddBlock(&blockchain.Block{})

	for _, i := range chain.Chain {
		fmt.Println(i)
	}
}
