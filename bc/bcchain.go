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
	}
	return true
}

func NewChain(initialBlock *Block) *Chain {
	chain := &Chain{}
	chain.Blocks = make([]Block, 0)
	chain.Blocks = append(chain.Blocks, *initialBlock)
	chain.LatestHash = initialBlock.Hash
	return chain
}
