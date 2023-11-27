package gaudius

import (
	"log"

	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AudiusSdkParams struct {
	DiscoveryNode        string
	CreatorNode          string
	EntityManagerAddress string
	AcdcGatewayRpc       string
}

type AudiusSdk struct {
	Discovery            *DiscoveryNode
	Storage              *StorageNode
	EntityManager        *contracts.EntityManager
	EntityManagerAddress *common.Address
	AcdcClient           *ethclient.Client
}

func NewAudiusSdkMainnetParams() *AudiusSdkParams {
	return &AudiusSdkParams{
		DiscoveryNode:        "https://discoveryprovider3.audius.co",
		CreatorNode:          "https://creatornode3.audius.co",
		EntityManagerAddress: "0x1Cd8a543596D499B9b6E7a6eC15ECd2B7857Fd64",
		AcdcGatewayRpc:       "https://acdc-gateway.audius.co/",
	}
}

func NewAudiusSdkTestnetParams() *AudiusSdkParams {
	return &AudiusSdkParams{
		DiscoveryNode:        "https://discoveryprovider2.staging.audius.co/",
		CreatorNode:          "https://creatornode5.staging.audius.co/",
		EntityManagerAddress: "0x1Cd8a543596D499B9b6E7a6eC15ECd2B7857Fd64",
		AcdcGatewayRpc:       "https://acdc-gateway.staging.audius.co/",
	}
}

func NewCustomSdk(params *AudiusSdkParams) (*AudiusSdk, error) {
	discovery := NewDiscoveryNode(params.DiscoveryNode)
	storage := NewStorageNode(params.CreatorNode)

	addr := common.HexToAddress(params.EntityManagerAddress)
	cl, err := ethclient.Dial(params.AcdcGatewayRpc)

	if err != nil {
		return nil, err
	}

	em, err := contracts.NewEntityManager(addr, cl)
	if err != nil {
		return nil, err
	}

	return &AudiusSdk{Discovery: discovery, Storage: storage, AcdcClient: cl, EntityManager: em, EntityManagerAddress: &addr}, nil
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
