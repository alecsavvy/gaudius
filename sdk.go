package gaudius

import (
	"log"

	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AudiusSdk struct {
	Discovery *DiscoveryNode
	Storage *StorageNode
	EntityManager *contracts.EntityManager
	AcdcClient *ethclient.Client
}

// initializes a new sdk against mainnet
func NewSdk() (*AudiusSdk, error) {
	discovery := NewDiscoveryNode("https://discoveryprovider3.audius.co")
	storage := NewStorageNode("https://creatornode.audius.co")

	addr := common.HexToAddress("0x1Cd8a543596D499B9b6E7a6eC15ECd2B7857Fd64")
	cl, err := ethclient.Dial("https://acdc-gateway.audius.co/")

	if err != nil {
		return nil, err
	}

	em, err := contracts.NewEntityManager(addr, cl)
	if err != nil {
		return nil, err
	}

	return &AudiusSdk{ Discovery: discovery, Storage: storage, AcdcClient: cl, EntityManager: em }, nil
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
