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
)

const (
	HTTPsProvider   = "wss://mainnet.infura.io/ws/v3/b754952dfaa24260ac2777131f6acd9b"
	ContractAddress = "0x1554F51F18F8E3fBe83E4442420E40Efc57ff446"
)

func main() {
	client, err := ethclient.Dial(HTTPsProvider)
	if err != nil {
		log.Fatalln(err)
	}
	contract, err := abibindings.NewContractabi(common.HexToAddress(ContractAddress), client)
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
		case <-headers:
			go func() {
				pausedChan := make(chan bool)
				whitelistChan := make(chan bool)
				go func() {
					p, err := contract.Paused(&bind.CallOpts{})
					if err != nil {
						log.Fatalln(err)
					}
					pausedChan <- p
				}()
				go func() {
					wl, err := contract.WhitelistActive(&bind.CallOpts{})
					if err != nil {
						log.Fatalln(err)
					}
					whitelistChan <- wl
				}()
				paused := <-pausedChan
				whitelist := <-whitelistChan
				if !paused && !whitelist {
					fmt.Println("MINT!")
				}
			}()
		}
	}

}
