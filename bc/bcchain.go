package bc

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

func NewChain(initialBlock *Block) *Chain {
	chain := &Chain{}
	chain.Blocks = make([]Block, 0)
	chain.Blocks = append(chain.Blocks, *initialBlock)
	chain.LatestHash = initialBlock.Hash
	return chain
}
