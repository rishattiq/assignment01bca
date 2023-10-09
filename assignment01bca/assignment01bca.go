package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

// the entire blockchain.
type Blockchain struct {
	Blocks []*Block
}

// New block is created and added to blockchain.
func (bc *Blockchain) NewBlock(transaction string, nonce int, previousHash string) {
	newBlock := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	newBlock.CurrentHash = bc.CreateHash(newBlock)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// Hash created for given block.
func (bc *Blockchain) CreateHash(block *Block) string {
	value := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	hash_value := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", hash_value)
}

// prints all the blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Transaction: %s\n Nonce: %d\n Previous Hash: %s\n Current Hash: %s\n\n",
			block.Transaction, block.Nonce, block.PreviousHash, block.CurrentHash)
	}
}

// the transaction of a given block is changed.
func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(bc.Blocks) {
		bc.Blocks[blockIndex].Transaction = newTransaction
		bc.Blocks[blockIndex].CurrentHash = bc.CreateHash(bc.Blocks[blockIndex])
	}
}

// verifies the integrity of the blockchain.
func (bc *Blockchain) VerifyChain() bool {
	for j := 1; j < len(bc.Blocks); j++ {
		if bc.Blocks[j].PreviousHash != bc.Blocks[j-1].CurrentHash {
			return false
		}
	}
	return true
}

// func main() {
// 	// Create a new blockchain
// 	bc := &Blockchain{}

// 	// Create the first block (Genesis block)
// 	genesisBlock := &Block{
// 		Transaction:  "Transaction 1",
// 		Nonce:        10000,
// 		PreviousHash: "GenesisBlockHash",
// 	}
// 	genesisBlock.CurrentHash = bc.CreateHash(genesisBlock)

// 	// Add the Genesis block to the blockchain
// 	bc.Blocks = append(bc.Blocks, genesisBlock)

// 	// Create subsequent blocks
// 	fmt.Print("\nAdd new block\n")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	if scanner.Scan() {
// 		transaction := scanner.Text()
// 		bc.NewBlock(transaction, 4, bc.Blocks[len(bc.Blocks)-1].CurrentHash)
// 	}
// 	bc.NewBlock("Transaction 2", 1, bc.Blocks[len(bc.Blocks)-1].CurrentHash)
// 	bc.NewBlock("Transaction 3", 2, bc.Blocks[len(bc.Blocks)-1].CurrentHash)

// 	// Display the blockchain
// 	bc.DisplayBlocks()

// 	// Verify the integrity of the blockchain
// 	if bc.VerifyChain() {
// 		fmt.Println("Blockchain is valid.")
// 	} else {
// 		fmt.Println("Blockchain is not valid.")
// 	}
// }
