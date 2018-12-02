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
}
