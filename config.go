package gaudius

import (
	"fmt"

	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AudiusSdkParams struct {
	DiscoveryNodes       []string
	CreatorNodes         []string
	EntityManagerAddress string
	RegistryAddress      string
	EthereumRpc          string
}

func NewAudiusSdkMainnetParams() *AudiusSdkParams {
	return &AudiusSdkParams{
		DiscoveryNodes:       MainnetDiscoveryNodesCached,
		CreatorNodes:         MainnetStorageNodesCached,
		EntityManagerAddress: MainnetAcdcAddress,
		RegistryAddress:      MainnetRegistryAddress,
	}
}

func NewAudiusSdkTestnetParams() *AudiusSdkParams {
	return &AudiusSdkParams{
		DiscoveryNodes:       TestnetDiscoveryNodesCached,
		CreatorNodes:         TestnetStorageNodesCached,
		EntityManagerAddress: TestnetAcdcAddress,
		RegistryAddress:      TestnetRegistryAddress,
	}
}

func NewCustomSdk(params *AudiusSdkParams) (*AudiusSdk, error) {
	discovery, err := NewDiscoveryNode(params.DiscoveryNodes)
	if err != nil {
		return nil, err
	}
	storage, err := NewStorageNode(params.CreatorNodes)
	if err != nil {
		return nil, err
	}

	addr := common.HexToAddress(params.EntityManagerAddress)
	acdcEndpoint := fmt.Sprintf("%s/chain", discovery.SelectedNode)
	cl, err := ethclient.Dial(acdcEndpoint)

	if err != nil {
		return nil, err
	}

	em, err := contracts.NewEntityManager(addr, cl)
	if err != nil {
		return nil, err
	}

	sdk := &AudiusSdk{Discovery: discovery, Storage: storage, AcdcClient: cl, EntityManager: em, EntityManagerAddress: &addr}

	// conditional features

	// init eth rpc and contracts if eth rpc supplied
	if params.EthereumRpc != "" {
		ethrpc, err := ethclient.Dial(params.EthereumRpc)
		if err != nil {
			return nil, err
		}
		contracts, err := NewAudiusContracts(ethrpc, params.RegistryAddress)
		if err != nil {
			return nil, err
		}

		sdk.EthereumClient = ethrpc
		sdk.Contracts = contracts
	}

	return sdk, nil
}
