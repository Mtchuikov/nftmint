package main

import (
	"context"
	"find-contract/abibindings"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"sync"
)

const (
	HTTPsProvider   = "wss://ropsten.infura.io/ws/v3/"
	ContractAddress = "0x192e52E8f5BD2C28a011706624d006bae14bB6B2"
)

func main() {
	client, err := ethclient.Dial(HTTPsProvider)
	if err != nil {
		log.Fatalln(err)
	}
	contract, err := abibindings.NewContractabiCaller(common.HexToAddress(ContractAddress), client)
	if err != nil {
		log.Fatalln(err)
	}
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	defer sub.Unsubscribe()
	if err != nil {
		log.Fatalln(err)
	}
	for {
		select {
		case err = <-sub.Err():
			log.Fatalln(err)
		case header := <-headers:
			go func() {
				var wg sync.WaitGroup
				var paused, whitelist bool
				go func() {
					wg.Add(1)
					p, err := contract.Paused(&bind.CallOpts{BlockNumber: header.Number})
					if err != nil {
						log.Fatalln(err)
					}
					paused = p
					wg.Done()
				}()
				go func() {
					wg.Add(1)
					wl, err := contract.WhitelistActive(&bind.CallOpts{BlockNumber: header.Number})
					if err != nil {
						log.Fatalln(err)
					}
					whitelist = wl
					wg.Done()
				}()
				wg.Wait()
				fmt.Println(paused, whitelist)
				if !paused && !whitelist {
					fmt.Println(1)
					return
				}
				fmt.Println(2)
			}()
		}
	}
}
