package main

import (
	"fmt"
	"github.com/guildenstern70/blockchain/bc"
	"time"
)

// VERSION OF GO-BLOCKCHAIN
const VERSION = "0.1"

func main() {
	fmt.Printf("GoBLOCKCHAIN v.%s\n", VERSION)

	// Step 1. Create Data
	var cars = createData()

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
	blockChain := bc.NewChain(initialBlock)

	for index := 1; index < len(cars); index++ {
		newBlock := bc.NewCarBlock(blockChain.LatestHash, *cars[index])
		blockChain.Append(newBlock)
	}

	return blockChain

}
