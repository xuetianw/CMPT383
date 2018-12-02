package blockchain

import (
	"work_queue"
)

type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	Start uint64
	End   uint64
	Blk   Block
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

func (mw miningWorker) Run() interface{} {
	result := new(MiningResult)

	for i := mw.Start; i <= mw.End; i++ {
		mw.Blk.Proof = i
		mw.Blk.Hash = mw.Blk.CalcHash()
		if mw.Blk.ValidHash() {
			result.Proof = i
			result.Found = true
			return result
		}
	}

	result.Found = false

	return result
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	// TODO
	chunkSize := (end - start) / chunks

	q := work_queue.Create(uint(workers), uint(chunks))
	for i := uint64(0); i+chunkSize <= end; i = i + chunkSize + 1 {
		mw := new(miningWorker)
		mw.Start = i
		mw.End = func() uint64 {
			if i+chunkSize > end {
				return end
			} else {
				return i + chunkSize
			}
		}()

		mw.Blk = blk
		q.Enqueue(mw)
	}

	var mr *MiningResult

	for r := range q.Results {
		mr = r.(*MiningResult)
		if mr.Found {
			q.Shutdown()
			break
		}
	}

	return *mr
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << (8 * blk.Difficulty)) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}
