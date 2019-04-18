package bc

import (
	"bytes"
	"fmt"
	"strconv"
)

// Blockchain chain
type Chain struct {
	Blocks     []Block
	LatestHash []byte
	Difficulty int
}

// Return the blockchain len
func (c *Chain) Append(block *Block) int {
	c.Blocks = append(c.Blocks, *block)
	c.LatestHash = block.Hash
	return len(c.Blocks)
}

// Return chain validity
func (c *Chain) IsValid() bool {

	for b := range c.Blocks {
		blockHash := PrintBytes(c.Blocks[b].Hash)

		if !bytes.Equal(c.Blocks[b].Hash, c.Blocks[b].ComputeHash()) {
			fmt.Println("Current Hashes not equal [Block #" + strconv.Itoa(b) + "]")
			return false
		}
		if b > 1 {
			if !bytes.Equal(c.Blocks[b-1].Hash, c.Blocks[b].PreviousHash) {
				fmt.Println("Previous Hashes not equal [Block #" + strconv.Itoa(b) + "]")
				return false
			}
		}
		if blockHash[0:c.Difficulty] != TargetDifficulty(c.Difficulty) {
			fmt.Println("Block has not been mined [Block #" + strconv.Itoa(b) + "]")
			return false
		}
	}
	return true
}

func NewChain(initialBlock *Block, difficulty int) *Chain {
	chain := &Chain{}
	chain.Difficulty = difficulty
	chain.Blocks = make([]Block, 0)
	chain.Blocks = append(chain.Blocks, *initialBlock)
	chain.LatestHash = initialBlock.Hash
	return chain
}
