package main

import (
	"context"
	"log"
	"math/big"
	"tz/client"
)

const blocksCount = 100

func main() {
	ctx := context.Background()
	c, err := client.NewClient()
	if err != nil {
		log.Panicf("can't initialize client with err %v", err)
	}
	log.Printf("fetching last %d blocks \n", blocksCount)
	blocks, err := c.FindBlocks(ctx, blocksCount)
	if err != nil {
		log.Panicf("can't fetch last 100 blocks, err - %v", err)
	}
	log.Printf("got %d blocks \n", blocksCount)
	log.Println("evaluating addresses deltas")
	addressesWithDelta := make(map[string]*big.Int)
	for _, block := range blocks {
		for _, transaction := range block.Transactions {
			transactionValue := stringByteToInt(transaction.Value)
			if transaction.To != nil {
				if _, hasAddressTo := addressesWithDelta[*transaction.To]; !hasAddressTo {
					addressesWithDelta[*transaction.To] = new(big.Int)
					_, setted := addressesWithDelta[*transaction.To].SetString("0", 16)
					if !setted {
						log.Panic("can't init value")
					}
				}
				addressesWithDelta[*transaction.To] = addressesWithDelta[*transaction.To].Add(addressesWithDelta[*transaction.To], transactionValue)
			}
			if _, hasAddressFrom := addressesWithDelta[transaction.From]; !hasAddressFrom {
				addressesWithDelta[transaction.From] = new(big.Int)
				_, setted := addressesWithDelta[transaction.From].SetString("0", 16)
				if !setted {
					log.Panic("can't init value")
				}
			}
			//если нужно посчитать сколько всего
			addressesWithDelta[transaction.From] = addressesWithDelta[transaction.From].Add(addressesWithDelta[transaction.From], transactionValue)
			//если нужно посчитать конечный баланс
			//addressesWithDelta[transaction.From] = addressesWithDelta[transaction.From].Sub(addressesWithDelta[transaction.From], transactionValue)
		}
	}
	var maxDelta *big.Int = big.NewInt(0)
	var addressMaxDelta string
	for address, delta := range addressesWithDelta {
		if delta.CmpAbs(maxDelta) == 1 {
			maxDelta = delta
			addressMaxDelta = address
		}
	}
	log.Printf("address - %s, delta - %s", addressMaxDelta, maxDelta.String())
}

func stringByteToInt(value string) *big.Int {
	n := new(big.Int)
	number, isSetted := n.SetString(value[2:], 16)
	if !isSetted {
		return nil
	}
	return number
}
