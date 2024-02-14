package gaudius

import (
	"log"

	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// the one struct to rule them all
type AudiusSdk struct {
	Discovery            *DiscoveryNode
	Storage              *StorageNode
	EntityManager        *contracts.EntityManager
	EntityManagerAddress *common.Address
	EthereumClient       *ethclient.Client
	AcdcClient           *ethclient.Client
	Oauth                *OauthConfiguration
	Contracts            *AudiusContracts
}

func NewTestnetSdk() (*AudiusSdk, error) {
	params := NewAudiusSdkTestnetParams()
	return NewCustomSdk(params)
}

func NewTestnetSdkUnsafe() *AudiusSdk {
	sdk, err := NewTestnetSdk()
	if err != nil {
		log.Fatal("failed to init testnet sdk: ", err)
	}
	return sdk
}

// initializes a new sdk against mainnet
func NewSdk() (*AudiusSdk, error) {
	params := NewAudiusSdkMainnetParams()
	return NewCustomSdk(params)
}

// initializes new sdk but panics on an error
// useful for testing or just quick scripts
// where specific error handling isn't needed
func NewSdkUnsafe() *AudiusSdk {
	sdk, err := NewSdk()
	if err != nil {
		log.Fatal("failed to init sdk: ", err)
	}
	return sdk
}
