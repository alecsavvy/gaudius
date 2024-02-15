package gaudius

import (
	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// the one struct to rule them all
type AudiusSdk struct {
	Discovery            *DiscoveryNode
	Content              *ContentNode
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

// initializes a new sdk against mainnet
func NewSdk() (*AudiusSdk, error) {
	params := NewAudiusSdkMainnetParams()
	return NewCustomSdk(params)
}
