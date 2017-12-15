// Package blockchain contains blocks and related methods
package blockchain

// GenesisBlock is the initial block
// Need to be filled with specific value
var GenesisBlock = Block{}

// BlockChain is exactly what the name suggests
// A chain of blocks
type BlockChain struct {
	Chain []*Block
}

// NewBlockChain generates a new Blockchain
func NewBlockChain() *BlockChain {
	return &BlockChain{Chain: []*Block{&GenesisBlock}}
}

// AddBlock adds a block to the chain
func (bc *BlockChain) AddBlock(b *Block) bool {
	// TODO: Check whether the newest block is valid or not
	// if not, don't update the chain and take necessary actions
	bc.Chain = append(bc.Chain, b)
	return true
}
