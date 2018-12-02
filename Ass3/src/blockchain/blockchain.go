package blockchain

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	// TODO
}

func (chain Blockchain) IsValid() bool {
	// TODO
	chainSize := len(chain.Chain)

	if chainSize == 0 {
		return true
	}

	blk := chain.Chain[0]

	// The initial block has previous hash all null bytes and is generation zero
	// And has valid hash
	if !byteArrayEqual(make([]byte, 32), blk.PrevHash) ||
		blk.Generation != 0 ||
		!byteArrayEqual(blk.Hash, blk.CalcHash()) ||
		!blk.ValidHash() {
		return false
	}

	// Valid hash from last to the initial block
	for i := chainSize - 1; i > 0; i-- {
		blkCur := chain.Chain[i]
		blkPrev := chain.Chain[i-1]

		// Each block has a generation value that is one more than the previous block.
		// Each block has the same difficulty value
		// Each block's previous hash matches the previous block's hash.
		// Each block's hash value actually matches its contents.
		// Each block's hash value ends in difficulty null bytes.
		if blkCur.Generation != blkPrev.Generation+1 ||
			blkCur.Difficulty != blkPrev.Difficulty ||
			!byteArrayEqual(blkCur.PrevHash, blkPrev.Hash) ||
			!byteArrayEqual(blkCur.Hash, blkCur.CalcHash()) ||
			!blkCur.ValidHash() {
			return false
		}
	}

	return true
}

func byteArrayEqual(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
