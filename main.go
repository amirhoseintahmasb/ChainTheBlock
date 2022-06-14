package main

import (
	"context"
	"fmt"

	"log"

	go1inch "github.com/jon4hz/go-1inch"
)

func main() {

	//create apis client
	
	client := go1inch.NewClient()
	//res, _, err := client.Healthcheck(context.Background(), "eth")
	// res, _, err := client.ApproveTransaction(context.Background(), "eth", "0x6b175474e89094c44da98b954eedeac495271d0f", &go1inch.ApproveTransactionOpts{
	// 	Amount: "5000000",
	// })
	res, _, err := client.Tokens(context.Background(), "eth")
	//call swapOpts
	//res, err := SwapOpts(context.Background(), "eth", "0x6b175474e89094c44da98b954eedeac495271d0f", "0x6b175474e89094c44da98b954eedeac495271d0f", "5000000", "50000")

	if err != nil {
		log.Fatal("Error while getting an approve calldata ", err)
	}
	fmt.Println(res)
}
