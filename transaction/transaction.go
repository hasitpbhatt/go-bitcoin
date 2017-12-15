package transaction

import "hash"

// Transaction contains the transaction details
type Transaction struct {
	Inputs  []*TxIn
	Outputs []*TxOut
}

// TxIn contains the input of transaction
type TxIn struct {
	// PreviousOutPoint contains the outpoint of previous transaction
	PreviousOutPoint *OutPoint

	// Need to add witness, signing script, and flag for lock time
	//
	// lock-time is the time before which the transaction won't be
	// considered valid
	// Reference: https://en.bitcoin.it/wiki/NLockTime
}

// TxOut contains the output of transaction
type TxOut struct {
	Amount int64

	// Script public key is needed along with the amount
}

// OutPoint contains the hash of the transaction and index
type OutPoint struct {
	// hash of transaction
	hash hash.Hash

	// index in the array of TxOut in Transaction
	index uint32
}
