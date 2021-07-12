package contracts

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

// Test inbox contract gets deployed correctly
func TestInitCascade(t *testing.T) {

	chainID := big.NewInt(1337)

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}
	blockGasLimit := uint64(4712388)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	//Deploy contract
	_, _, contract, _ := DeployCascade(auth, client)

	// commit all pending transactions
	client.Commit()

	// Set init message
	contract.SetMessage(auth, "Hello World")

	client.Commit()

	if got, _ := contract.Message(nil); got != "Hello World" {
		t.Errorf("Expected message to be: Hello World. Got: %s", got)
	}

	// Set other message
	contract.SetMessage(auth, "New message")

	client.Commit()

	if got, _ := contract.Message(nil); got != "New message" {
		t.Errorf("Expected message to be: New message. Got: %s", got)
	}
}
