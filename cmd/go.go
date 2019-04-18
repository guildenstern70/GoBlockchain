package main

import (
	"fmt"
	"github.com/guildenstern70/blockchain/bc"
	"time"
)

// VERSION OF GO-BLOCKCHAIN
const VERSION = "0.2"
const DIFFICULTY = 3

func main() {
	fmt.Printf("GoBLOCKCHAIN v.%s\n\n", VERSION)

	// Step 1. Create Data
	cars := createData()

	// Step 2: Print cars
	for c := range cars {
		fmt.Println(cars[c].ToString())
	}

	// Step 3: Create blockchain
	blockchain := createBlockChain(cars)

	// Step 4: Print blockchain
	for b := range blockchain.Blocks {
		fmt.Println(blockchain.Blocks[b].ToString())
	}

	// Step 5: Validate blockchain
	if blockchain.IsValid() {
		fmt.Println("Blockchain is VALID.")
	} else {
		fmt.Println("** Blockchain is NOT valid. **")
	}

}

func createData() []*bc.CarData {

	date1 := time.Date(2017, 12, 26, 10, 00, 00, 0, time.UTC)
	date2 := time.Date(2018, 6, 12, 11, 34, 00, 0, time.UTC)
	date3 := time.Date(2018, 12, 14, 17, 33, 00, 0, time.UTC)
	now := time.Now()

	carsCollection := make([]*bc.CarData, 0)
	carsCollection = append(carsCollection, bc.NewBcData(date1, "Alessio", 2000))
	carsCollection = append(carsCollection, bc.NewBcData(date2, "Rocco", 1800))
	carsCollection = append(carsCollection, bc.NewBcData(date3, "Giulia", 1500))
	carsCollection = append(carsCollection, bc.NewBcData(now, "Marina", 1000))

	return carsCollection

}

func createBlockChain(cars []*bc.CarData) *bc.Chain {

	initialBlock := bc.InitialCarBlock(*cars[0])
	initialBlock.Mine(DIFFICULTY)
	blockChain := bc.NewChain(initialBlock)

	for index := 1; index < len(cars); index++ {
		newBlock := bc.NewCarBlock(blockChain.LatestHash, *cars[index])
		newBlock.Mine(DIFFICULTY)
		blockChain.Append(newBlock)
	}

	return blockChain

}
