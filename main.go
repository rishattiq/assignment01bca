// package main

// import (
// 	"fmt"
// )

// func main() {
// 	bc := &blockchain.Blockchain{}

// 	// Add blocks to the blockchain
// 	bc.NewBlock("Alice to Bob", 12345, "0")
// 	bc.NewBlock("Bob to Charlie", 67890, bc.Blocks[len(bc.Blocks)-1].CurrentHash)

// 	// Display all blocks
// 	fmt.Println("All Blocks:")
// 	bc.DisplayBlocks()

// 	// Change the transaction of a block
// 	bc.ChangeBlock(1, "Updated transaction")

// 	// Verify the blockchain
// 	if bc.VerifyChain() {
// 		fmt.Println("Blockchain is valid.")
// 	} else {
// 		fmt.Println("Blockchain is not valid.")
// 	}
// }
