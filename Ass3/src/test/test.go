package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Block struct {
	Generation uint64
	Difficulty uint8
	Data       string
	PrevHash   []byte
	Hash       []byte
	Proof      uint64
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	// TODO
	initialBlock := new(Block)
	initialBlock.Generation = 0
	initialBlock.Difficulty = difficulty
	initialBlock.Data = ""
	initialBlock.PrevHash = make([]byte, 32)
	return *initialBlock
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO
	nextBlock := new(Block)
	nextBlock.Generation = prev_block.Generation + 1
	nextBlock.Difficulty = prev_block.Difficulty
	nextBlock.Data = data
	nextBlock.PrevHash = prev_block.Hash
	return *nextBlock
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	// TODO
	value := hex.EncodeToString(blk.PrevHash) +
		":" + strconv.FormatUint(blk.Generation, 10) +
		":" + strconv.FormatUint(uint64(blk.Difficulty), 10) +
		":" + blk.Data +
		":" + strconv.FormatUint(blk.Proof, 10)

	data := []byte(value)
	arr := sha256.Sum256(data)
	return arr[:]
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	hash := blk.CalcHash()
	length := len(hash)
	if length != len(blk.Hash) {
		return false
	}
	for i := length - int(blk.Difficulty); i < length; i++ {
		if hash[i] != byte(0) {
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

func main() {

	zeros := make([]byte, 32)

	var c = hex.EncodeToString(zeros)
	fmt.Println(c)

	value := hex.EncodeToString(make([]byte, 32)) +
		":" + strconv.FormatUint(0, 10) +
		":" + strconv.FormatUint(uint64(16), 10) +
		":" + "" +
		":" + strconv.FormatUint(56231, 10)

	fmt.Println(value)

	fmt.Println(strconv.FormatUint(10, 10))

	b0 := Initial(16)
	//b0.Mine(1)
	b1 := b0.Next("message")
	//b1.Mine(1)
	println(b0.CalcHash())
	//    b1 := b0.Next("message")
	//println(b0.ValidHash())

	fmt.Println(hex.EncodeToString(b0.CalcHash()))
	fmt.Println(hex.EncodeToString(b1.CalcHash()))
}
