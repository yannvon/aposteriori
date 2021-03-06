package contracts

import (
	"context"
	"log"
	"math/big"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Test inbox contract gets deployed correctly
func TestMultishotEvents(t *testing.T) {

	chainID := big.NewInt(1337)

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	privateKey2, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	auth2, _ := bind.NewKeyedTransactorWithChainID(privateKey2, chainID)

	privateKey3, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	auth3, _ := bind.NewKeyedTransactorWithChainID(privateKey3, chainID)

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
		auth2.From: {
			Balance: balance,
		},
		auth3.From: {
			Balance: balance,
		},
	}
	blockGasLimit := uint64(4712388)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	//Deploy contract
	validators := []common.Address{auth.From, auth2.From, auth3.From}
	stake := []*big.Int{balance, balance, balance}

	contractAddress, _, contract, _ := DeployMultishot(auth, client, validators, stake)

	// Handle events
	go handleEvents(client, t, contractAddress, contract)

	// commit all pending transactions
	client.Commit()

	// Prepare test
	txSender := auth.From
	txNonce := big.NewInt(1)
	txHash := big.NewInt(42)

	// Read if decision take already
	if got, _ := contract.Read(nil, txSender, txNonce); got.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected no decision. Got: %d", got)
	}

	// Propose for a specific instance
	var e error

	_, e = contract.Propose(auth2, txSender, txNonce, txHash)
	if e != nil {
		panic(e)
	}
	_, e = contract.Propose(auth2, txSender, txNonce, txHash)
	if e == nil {
		panic("Expected error to be thrown for duplicate propose.")
	}
	_, e = contract.Propose(auth3, txSender, txNonce, txHash)
	if e != nil {
		panic(e)
	}
	client.Commit()

	if got, _ := contract.Read(nil, txSender, txNonce); got.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected no decision yet. Got: %d", got)
	}

	// Last propose should make 4/5 of stake
	contract.Propose(auth, txSender, txNonce, txHash)

	client.Commit()

	// t.Log("txSender", txSender)
	if got, _ := contract.Read(nil, txSender, txNonce); got.Cmp(txHash) != 0 {
		t.Errorf("Expected decision to be sole proposed txHash. Got: %d", got)
	}

}

func handleEvents(client *backends.SimulatedBackend, t *testing.T, contractAddress common.Address, contract *Multishot) {
	// Subscribe to events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			// Example how the log can be parsed back to object and relevant fields extracted
			decided, _ := contract.ParseDecided(vLog)
			t.Log(decided.Decision, decided.TxOrigin, decided.TxNonce)
		}
	}
}
