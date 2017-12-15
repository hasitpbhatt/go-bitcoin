package blockchain

import "hash"

// Header contains information about the block e.g.
// version, timestamp, previous hash, merkle root etc.
type Header struct {
	// timestamp is approximate unix timestamp of block creation
	timestamp uint32

	// difficulty indicates the proof-of-work algorithm difficulty target for this block
	// equivaluent to nBits in original bitcoin implementation
	difficulty uint32

	// version of the node
	version int32

	// PreviousHash is hash of previous block
	PreviousHash hash.Hash32

	// MerkleRoot is merkle root of current block
	MerkleRoot hash.Hash32
}
