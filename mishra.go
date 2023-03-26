package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/Polygon/matic-sdk-bindings/go-bindings/nft"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a wallet address as an argument.")
	}
	walletAddress := common.HexToAddress(os.Args[1])

	// Connect to Polygon Mumbai Testnet
	client, err := ethclient.Dial("https://rpc-mumbai.maticvigil.com")
	if err != nil {
		log.Fatal(err)
	}

	// Get the NFT contract instance
	contractAddress := common.HexToAddress("0x16581f93797e33fd2b1a3497822adf1762ee36e2")
	instance, err := nft.NewNft(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get the total number of tokens for the given address
	totalTokens, err := instance.BalanceOf(&bind.CallOpts{}, walletAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total number of tokens: %v\n", totalTokens)

	// Lookup the metadata for the first token (if available) and display
	if totalTokens.Cmp(big.NewInt(0)) > 0 {
		// Get the token ID of the first token
		tokenID := big.NewInt(0)

		// Get the metadata for the token
		tokenMetadata, err := instance.TokenURI(&bind.CallOpts{}, tokenID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Token metadata: %v\n", tokenMetadata)
	} else {
		fmt.Println("No tokens in wallet.")
	}
}
