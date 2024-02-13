package gaudius

import "errors"

var (
	Mainnet = "mainnet"
	Testnet = "testnet"
	Devnet  = "devnet"
	Prod    = "prod"
	Stage   = "staging"
	Local   = "local"
)

type AudiusEnvironment interface {
	GetDiscoveryEndpoints() ([]string, error)
	GetStorageEndpoints() ([]string, error)
	GetEnvironment() string
	GetContracts() (*AudiusContracts, error)
}

type MainnetEnvironment struct {
	cachedEndpoints    bool
	discoveryEndpoints []string
	storageEndpoints   []string
}

func (env *MainnetEnvironment) GetDiscoveryEndpoints() ([]string, error) {
	if len(env.discoveryEndpoints) > 0 {
		return env.discoveryEndpoints, nil
	}
	if env.cachedEndpoints {
		env.discoveryEndpoints = MainnetDiscoveryNodesCached
		return env.discoveryEndpoints, nil
	}
	// TODO: read all nodes from chain
	return []string{}, errors.New("no discovery nodes found")
}

func (env *MainnetEnvironment) GetStorageEndpoints() ([]string, error) {
	if len(env.storageEndpoints) > 0 {
		return env.storageEndpoints, nil
	}
	if env.cachedEndpoints {
		env.storageEndpoints = MainnetStorageNodesCached
		return env.storageEndpoints, nil
	}
	// TODO: read all nodes from chain
	return []string{}, errors.New("no storage nodes found")
}

func (env *MainnetEnvironment) GetEnvironment() string {
	return Mainnet
}

type TestnetEnvironment struct {
	cachedEndpoints    bool
	discoveryEndpoints []string
	storageEndpoints   []string
}

type DevnetEnvironment struct {
	discoveryEndpoints []string
	storageEndpoints   []string
}
