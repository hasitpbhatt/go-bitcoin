// Package blockchain contains blocks and related methods
package blockchain

import "github.com/hasitpbhatt/go-bitcoin/transaction"

// Block contains all information about the block
type Block struct {
	// Size of the block
	Size int32

	// Header contains all the information like merkle root, version, difficulty
	Header Header

	// TransactionCounter is number of transactions block contains
	TransactionCounter []byte

	// Transactions is array of transactions in block
	Transactions []*transaction.Transaction
}
