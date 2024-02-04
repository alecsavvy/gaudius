package gaudius

import (
	"errors"
	"fmt"

	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// manager struct so contracts get loaded lazily upon usage
// and don't spam rpcs for contracts a developer may not need
type AudiusContracts struct {
	Rpc                    *ethclient.Client
	Registry               *contracts.Registry
	Governance             *contracts.Governance
	Staking                *contracts.Staking
	ServiceProviderFactory *contracts.ServiceProviderFactory
	ClaimsManager          *contracts.ClaimsManager
	DelegateManager        *contracts.DelegateManager
	AudioToken             *contracts.AudiusToken
	RewardsManager         *contracts.EthRewardsManager
}

func NewAudiusContracts(rpc *ethclient.Client, registryAddress string) (*AudiusContracts, error) {
	ok := common.IsHexAddress(registryAddress)
	if !ok {
		return nil, errors.New(fmt.Sprintf("registryAddress %s is not a valid hex address", registryAddress))
	}
	addr := common.HexToAddress(registryAddress)
	registry, err := contracts.NewRegistry(addr, rpc)
	if err != nil {
		return nil, err
	}
	return &AudiusContracts{
		Rpc:      rpc,
		Registry: registry,
	}, nil
}

func (ac *AudiusContracts) GetAudioTokenContract() (*contracts.AudiusToken, error) {
	if ac.AudioToken != nil {
		return ac.AudioToken, nil
	}

	addr, err := ac.Registry.GetContract0(nil, AudiusTokenKey)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewAudiusToken(addr, ac.Rpc)
	if err != nil {
		return nil, err
	}

	ac.AudioToken = contract
	return contract, nil
}
