package bc

import (
	"fmt"
	"strconv"
	"time"
)

// Blockchain block
type Block struct {
	Hash         []byte
	PreviousHash []byte
	Data         CarData
	TimeStamp    time.Time
	Nonce        int
}

func (b *Block) Finalize() {
	b.Hash = b.ComputeHash()
}

func (b *Block) ComputeHash() []byte {
	return GetHash(b.hashText())
}

func (b *Block) Mine(difficulty int) {
	target := createTarget(difficulty)
	bHash := PrintBytes(b.Hash)
	for bHash[0:difficulty] != target {
		b.Nonce += 1
		b.Hash = b.ComputeHash()
		bHash = PrintBytes(b.Hash)
	}
	fmt.Printf("Block Mined: %s (nonce = %d)\n", bHash, b.Nonce)
}

func (b *Block) hashText() string {
	text := PrintBytes(b.PreviousHash)
	text += b.Data.Owner
	text += b.TimeStamp.Format(time.RFC3339)
	text += strconv.Itoa(b.Nonce)
	return text
}

func (b *Block) ToString() string {
	previousHash := PrintBytes(b.PreviousHash)
	hash := PrintBytes(b.Hash)
	carOwner := b.Data.Owner
	nonceString := strconv.Itoa(b.Nonce)
	blockString := "Block: \n"
	blockString += "PreviousHash => " + previousHash + "\n"
	blockString += "Hash         => " + hash + "\n"
	blockString += "Owner        => " + carOwner + "\n"
	blockString += "Nonce        => " + nonceString + "\n"
	return blockString
}

func NewCarBlock(previousHash []byte, carData CarData) *Block {
	b := &Block{}
	b.Data = carData
	b.PreviousHash = previousHash
	b.Nonce = 0
	b.Finalize()
	return b
}

func InitialCarBlock(carData CarData) *Block {
	b := &Block{}
	b.Data = carData
	b.PreviousHash = make([]byte, 0)
	b.Nonce = 0
	b.Finalize()
	return b
}

/** Privates **/

func createTarget(difficulty int) string {
	target := ""
	for x := 0; x < difficulty; x++ {
		target += "0"
	}
	return target
}
