package blockchain

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: some useful tests of Blocks

// Test that a initial block can be created successfully
func TestCreateInitialBlock(t *testing.T) {
	assertUtil := assert.New(t)
	blk := Initial(4)

	assertUtil.Equal(uint64(0), blk.Generation)
	assertUtil.Equal("", blk.Data)
	assertUtil.Equal(uint8(4), blk.Difficulty)

	for _, i := range blk.PrevHash {
		assertUtil.Equal(byte(0), i)
	}
}

// Test that hash can be calculated correctly
func TestCalcHash(t *testing.T) {
	assertUtil := assert.New(t)
	b0 := Initial(2)
	b0.Proof = 242278

	b0hash := hex.EncodeToString(b0.CalcHash())
	assertUtil.Equal("29528aaf90e167b2dc248587718caab237a81fd25619a5b18be4986f75f30000", b0hash)
}

// Test that valid hash function works
func TestValidHash(t *testing.T) {
	assertUtil := assert.New(t)
	b0 := Initial(2)
	b0.Hash = b0.CalcHash()

	// Proof 0 is invalid
	assertUtil.False(b0.ValidHash())

	// Proof 242278 is valid
	b0.Proof = 242278
	b0.Hash = b0.CalcHash()
	assertUtil.True(b0.ValidHash())
}

// Test that mining with difficulty = 2 can find correct proof
func TestMiningWithDifficulty2(t *testing.T) {
	assertUtil := assert.New(t)
	b0 := Initial(2)
	b0.Mine(1)

	// Test b0 proof and hash
	assertUtil.Equal(uint64(242278), b0.Proof)
	assertUtil.Equal("29528aaf90e167b2dc248587718caab237a81fd25619a5b18be4986f75f30000",
		hex.EncodeToString(b0.CalcHash()))

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)

	// Test b1 proof and hash
	assertUtil.Equal(uint64(41401), b1.Proof)
	assertUtil.Equal("d558f4b9a0b5df021a98066efa5779758cd02f721ebcd8112872265799580000",
		hex.EncodeToString(b1.CalcHash()))

	b2 := b1.Next("this is not interesting")
	b2.Mine(1)

	// Test b2 proof and hash
	assertUtil.Equal(uint64(195955), b2.Proof)
	assertUtil.Equal("b611d6fc52964c1c89c717731807c5785ca6bf50d0922b0fc2fdf283757c0000",
		hex.EncodeToString(b2.CalcHash()))
}

// Test that mining with difficulty = 3 can find correct proof
func TestMiningWithDifficulty3(t *testing.T) {
	assertUtil := assert.New(t)
	b0 := Initial(3)
	b0.Mine(1)

	// Test b0 proof and hash
	assertUtil.Equal(uint64(8816998), b0.Proof)
	assertUtil.Equal("b7eeb8d8af2133911e09e5687c7b8da31807579fa3661e7dff6fa260e8000000",
		hex.EncodeToString(b0.CalcHash()))

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)

	// Test b1 proof and hash
	assertUtil.Equal(uint64(16634616), b1.Proof)
	assertUtil.Equal("196d83601de2eedd2d2dd6e10ae804c24f7d45297f16c293ed2c50cfcd000000",
		hex.EncodeToString(b1.CalcHash()))

	b2 := b1.Next("this is not interesting")
	b2.Mine(1)

	// Test b2 proof and hash
	assertUtil.Equal(uint64(8148543), b2.Proof)
	assertUtil.Equal("e91c5fdf5154cdaa0b3063fb4f9189beb43fd7e67d1c32e9f7e0aa2cd0000000",
		hex.EncodeToString(b2.CalcHash()))
}

func TestIsValid(t *testing.T) {
	assertUtil := assert.New(t)

	chain := new(Blockchain)

	// Test empty chain
	assertUtil.True(chain.IsValid())

	b0 := Initial(2)
	b0.Mine(1)
	chain.Add(b0)

	// Test if there only one block
	assertUtil.True(chain.IsValid())

	b1 := b0.Next("second block")
	b1.Mine(1)
	chain.Add(b1)
	b2 := b1.Next("third block")
	b2.Mine(1)
	chain.Add(b2)

	assertUtil.True(chain.IsValid())

	blkInvalid := Initial(2)
	blkInvalid.Mine(1)

	chain.Add(blkInvalid)
	assertUtil.False(chain.IsValid())
}
