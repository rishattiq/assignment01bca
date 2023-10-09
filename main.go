package main

import (
	"ass_01/assignment01bca"
	"bufio"
	"fmt"
	"os"
)

func main() {

	bc := &assignment01bca.Blockchain{}
	// Create the first block (Genesis block)
	genesisBlock := &assignment01bca.Block{
		Transaction:  "Transaction 1",
		Nonce:        10000,
		PreviousHash: "GenesisBlockHash",
	}
	genesisBlock.CurrentHash = bc.CreateHash(genesisBlock)

	// Add the Genesis block to the blockchain
	bc.Blocks = append(bc.Blocks, genesisBlock)

	// Create subsequent blocks
	fmt.Print("\nAdd new block\n")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		transaction := scanner.Text()
		bc.NewBlock(transaction, 4, bc.Blocks[len(bc.Blocks)-1].CurrentHash)
	}
	bc.NewBlock("Transaction 2", 1, bc.Blocks[len(bc.Blocks)-1].CurrentHash)
	bc.NewBlock("Transaction 3", 2, bc.Blocks[len(bc.Blocks)-1].CurrentHash)

	// Display the blockchain
	bc.DisplayBlocks()

	// Verify the integrity of the blockchain
	if bc.VerifyChain() {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
