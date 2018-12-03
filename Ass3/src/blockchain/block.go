package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	// TODO
	firstBlock := Block{
		Generation: 0,
		Difficulty: difficulty,
		Data:       "",
		PrevHash:   make([]byte, 32),
		//Hash: []byte(hex.EncodeToString([]byte(""))),

	}
	return firstBlock
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO
	newBlock := Block{
		PrevHash: prev_block.Hash,

		Generation: (prev_block.Generation + 1),
		Difficulty: (prev_block.Difficulty),
		Data:       data,
		//Hash: []byte(hex.EncodeToString(data)),

	}
	return newBlock

}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	// TODO

	value := hex.EncodeToString(blk.PrevHash) +
		":" + strconv.Itoa(int(blk.Generation)) +
		":" + strconv.Itoa(int(blk.Difficulty)) +
		":" + blk.Data +
		":" + strconv.Itoa(int(blk.Proof))

	data := []byte(value)

	h := sha256.New()
	h.Write([]byte(data))

	return h.Sum(nil)
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	// TODO

	hash := blk.CalcHash()
	length := len(hash)
	if len(blk.CalcHash()) != len(blk.Hash) {
		return false
	}

	lenthof_zero := blk.Difficulty
	for i := length - int(lenthof_zero); i < length; i++ {
		if hash[i] != 0 {
			return false
		}
	}

	for i := 0; i < length; i++ {
		if hash[i] != blk.Hash[i] {
			return false
		}
	}

	return true

}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
