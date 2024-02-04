package gaudius

import (
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

func NewAudiusSdkMainnetParams() *AudiusSdkParams {
	return &AudiusSdkParams{
		DiscoveryNode:        "https://discoveryprovider3.audius.co",
		CreatorNode:          "https://creatornode3.audius.co",
		EntityManagerAddress: MainnetAcdcAddress,
		AcdcGatewayRpc:       "https://acdc-gateway.audius.co/",
	}
}

func NewAudiusSdkTestnetParams() *AudiusSdkParams {
	return &AudiusSdkParams{
		DiscoveryNode:        "https://discoveryprovider2.staging.audius.co/",
		CreatorNode:          "https://creatornode5.staging.audius.co/",
		EntityManagerAddress: TestnetAcdcAddress,
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
