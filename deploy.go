package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yannvon/aposteriori/contracts"
)

func deploy() {
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/f9b32e6b21e740eab75d12e2e0318f3d")
	chainID := big.NewInt(4)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth1, err := bind.NewTransactorWithChainID(strings.NewReader(Key1), "", chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth2, err := bind.NewTransactorWithChainID(strings.NewReader(Key2), "", chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth3, err := bind.NewTransactorWithChainID(strings.NewReader(Key3), "", chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth4, err := bind.NewTransactorWithChainID(strings.NewReader(Key4), "", chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth5, err := bind.NewTransactorWithChainID(strings.NewReader(Key5), "", chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Set up contract init
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	validators := []common.Address{auth1.From, auth2.From, auth3.From, auth4.From, auth5.From}
	stake := []*big.Int{balance, balance, balance, balance, balance}

	// Deploy contract
	address, _, _, _ := contracts.DeployMultishot(auth1, blockchain, validators, stake)

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
}
