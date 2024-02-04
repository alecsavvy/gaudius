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
	ServiceTypeManager     *contracts.ServiceTypeManager
}

// instantiates audius contract manager so that contracts can be initialized lazily
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

// instantiates audius contract manager and eagerly gathers contract instances as well
func NewAudiusContractsInit(rpc *ethclient.Client, registryAddress string) (*AudiusContracts, error) {
	ac, err := NewAudiusContracts(rpc, registryAddress)
	if err != nil {
		return nil, err
	}

	// init all contracts
	_, err = ac.GetAudioTokenContract()
	if err != nil {
		return nil, errors.Join(errors.New("init audio contract failed"), err)
	}
	_, err = ac.GetGovernanceContract()
	if err != nil {
		return nil, errors.Join(errors.New("init governance contract failed"), err)
	}
	_, err = ac.GetStakingContract()
	if err != nil {
		return nil, errors.Join(errors.New("init staking contract failed"), err)
	}
	_, err = ac.GetServiceProviderFactoryContract()
	if err != nil {
		return nil, errors.Join(errors.New("init service provider factory contract failed"), err)
	}
	_, err = ac.GetClaimsManagerContract()
	if err != nil {
		return nil, errors.Join(errors.New("init claims manager contract failed"), err)
	}
	_, err = ac.GetDelegateManagerContract()
	if err != nil {
		return nil, errors.Join(errors.New("init delegate manager contract failed"), err)
	}
	_, err = ac.GetRewardsManagerContract()
	if err != nil {
		return nil, errors.Join(errors.New("init rewards manager contract failed"), err)
	}
	_, err = ac.GetServiceTypeManagerContract()
	if err != nil {
		return nil, errors.Join(errors.New("init service type manager contract failed"), err)
	}

	return ac, nil
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

func (ac *AudiusContracts) GetGovernanceContract() (*contracts.Governance, error) {
	if ac.Governance != nil {
		return ac.Governance, nil
	}

	addr, err := ac.Registry.GetContract0(nil, GovernanceKey)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewGovernance(addr, ac.Rpc)
	if err != nil {
		return nil, err
	}

	ac.Governance = contract
	return contract, nil
}

func (ac *AudiusContracts) GetStakingContract() (*contracts.Staking, error) {
	if ac.Staking != nil {
		return ac.Staking, nil
	}

	addr, err := ac.Registry.GetContract0(nil, StakingKey)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewStaking(addr, ac.Rpc)
	if err != nil {
		return nil, err
	}

	ac.Staking = contract
	return contract, nil
}

func (ac *AudiusContracts) GetServiceProviderFactoryContract() (*contracts.ServiceProviderFactory, error) {
	if ac.ServiceProviderFactory != nil {
		return ac.ServiceProviderFactory, nil
	}

	addr, err := ac.Registry.GetContract0(nil, ServiceProviderFactoryKey)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewServiceProviderFactory(addr, ac.Rpc)
	if err != nil {
		return nil, err
	}

	ac.ServiceProviderFactory = contract
	return contract, nil
}

func (ac *AudiusContracts) GetClaimsManagerContract() (*contracts.ClaimsManager, error) {
	if ac.ClaimsManager != nil {
		return ac.ClaimsManager, nil
	}

	addr, err := ac.Registry.GetContract0(nil, ClaimsManagerKey)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewClaimsManager(addr, ac.Rpc)
	if err != nil {
		return nil, err
	}

	ac.ClaimsManager = contract
	return contract, nil
}

func (ac *AudiusContracts) GetDelegateManagerContract() (*contracts.DelegateManager, error) {
	if ac.DelegateManager != nil {
		return ac.DelegateManager, nil
	}

	addr, err := ac.Registry.GetContract0(nil, DelegateManagerKey)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewDelegateManager(addr, ac.Rpc)
	if err != nil {
		return nil, err
	}

	ac.DelegateManager = contract
	return contract, nil
}

func (ac *AudiusContracts) GetRewardsManagerContract() (*contracts.EthRewardsManager, error) {
	if ac.RewardsManager != nil {
		return ac.RewardsManager, nil
	}

	addr, err := ac.Registry.GetContract0(nil, RewardsManagerKey)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewEthRewardsManager(addr, ac.Rpc)
	if err != nil {
		return nil, err
	}

	ac.RewardsManager = contract
	return contract, nil
}

func (ac *AudiusContracts) GetServiceTypeManagerContract() (*contracts.ServiceTypeManager, error) {
	if ac.ServiceTypeManager != nil {
		return ac.ServiceTypeManager, nil
	}

	spf, err := ac.GetServiceProviderFactoryContract()
	if err != nil {
		return nil, err
	}

	addr, err := spf.GetServiceTypeManagerAddress(nil)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewServiceTypeManager(addr, ac.Rpc)
	if err != nil {
		return nil, err
	}

	ac.ServiceTypeManager = contract
	return contract, nil
}
