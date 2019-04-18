package bc

import (
	"time"
)

// Blockchain block
type Block struct {
	Hash         []byte
	PreviousHash []byte
	Data         CarData
	TimeStamp    time.Time
}

func (b *Block) ComputeHash() []byte {
	return GetCarHash(b.Data)
}

func (b *Block) ToString() string {
	var previousHash = PrintBytes(b.PreviousHash)
	var hash = PrintBytes(b.Hash)
	var carOwner = b.Data.Owner
	var blockString = "Block: \n"
	blockString += "PreviousHash => " + previousHash + "\n"
	blockString += "Hash         => " + hash + "\n"
	blockString += "Owner        => " + carOwner + "\n"
	return blockString
}

func NewCarBlock(previousHash []byte, carData CarData) *Block {
	b := &Block{}
	b.Data = carData
	b.Hash = GetCarHash(carData)
	b.PreviousHash = previousHash
	return b
}

func InitialCarBlock(carData CarData) *Block {
	b := &Block{}
	b.Data = carData
	b.Hash = GetCarHash(carData)
	b.PreviousHash = make([]byte, 0)
	return b
}
