package gaudius

import (
	"fmt"
	"os"

	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
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
	discovery, err := NewDiscoveryNode([]string{})
	if err != nil {
		return nil, err
	}
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

// function just to sandbox the config flow
func ScratchConfigImplementation() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	envKey := "ENVIRONMENT"
	rpcKey := "ETHEREUM_RPC_URL"
	apiKeyKey := "AUDIUS_API_KEY"
	apiSecretKey := "AUDIUS_PRIVATE_KEY"

	environment := os.Getenv(envKey)
	ethereumRpc := os.Getenv(rpcKey)
	apiKey := os.Getenv(apiKeyKey)
	apiSecret := os.Getenv(apiSecretKey)

	if ethereumRpc == "" {
		// contracts are disabled
		// use cached nodes are forced
	}

	if environment == "" {
		environment = "mainnet"
	}

	// build relevant environment struct
	// that impls AudiusEnvironment interface

	if apiKey == "" {
		// oauth won't work
		// can't derive app name
	}

	if apiSecret == "" {
		// oauth won't work
		// writes disabled
	}

	eth, err := ethclient.Dial(ethereumRpc)
	if err != nil {
		return err
	}

	// determine discovery and storage nodes

	// determine selected discovery and storage nodes
	discprov := ""
	storage := ""

	// derive acdc rpc from discovery node
	acdc, err := ethclient.Dial(fmt.Sprintf("%s/chain", discprov))
	if err != nil {
		return err
	}

	// derive relay endpoint from discovery node
	relay := fmt.Sprintf("%s/relay", discprov)

	// bundle into AudiusSdk struct
	// with feature flag options

	return nil
}
