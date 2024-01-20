EXAMPLE?=get-user-and-image

all: format build test

deps:
	brew tap ethereum/ethereum
	brew install ethereum

generate-api:
	swagger generate model --skip-validation -f https://discoveryprovider3.audius.co/v1/swagger.json -m models -t ./gen/discovery
	swagger generate model --skip-validation -f https://discoveryprovider3.audius.co/v1/full/swagger.json -m models -t ./gen/discoveryFull
	go mod tidy

generate-contracts:
	# entity manager code generation
	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/data-contracts/ABIs/EntityManager.json | jq '.abi' > ./gen/contracts/EntityManager.json
	abigen --abi ./gen/contracts/EntityManager.json --pkg contracts --type EntityManager --out ./gen/contracts/entity_manager.go

	# eth contract code generation
	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/eth-contracts/ABIs/Registry.json | jq '.abi' > ./gen/contracts/Registry.json
	abigen --abi ./gen/contracts/Registry.json --pkg contracts --type Registry --out ./gen/contracts/registry.go

	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/eth-contracts/ABIs/Governance.json | jq '.abi' > ./gen/contracts/Governance.json
	abigen --abi ./gen/contracts/Governance.json --pkg contracts --type Governance --out ./gen/contracts/governance.go

	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/eth-contracts/ABIs/Staking.json | jq '.abi' > ./gen/contracts/Staking.json
	abigen --abi ./gen/contracts/Staking.json --pkg contracts --type Staking --out ./gen/contracts/staking.go

	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/eth-contracts/ABIs/ServiceProviderFactory.json | jq '.abi' > ./gen/contracts/ServiceProviderFactory.json
	abigen --abi ./gen/contracts/ServiceProviderFactory.json --pkg contracts --type ServiceProviderFactory --out ./gen/contracts/service_provider_factory.go

	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/eth-contracts/ABIs/ServiceTypeManager.json | jq '.abi' > ./gen/contracts/ServiceTypeManager.json
	abigen --abi ./gen/contracts/ServiceTypeManager.json --pkg contracts --type ServiceTypeManager --out ./gen/contracts/service_type_manager.go

	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/eth-contracts/ABIs/ClaimsManager.json | jq '.abi' > ./gen/contracts/ClaimsManager.json
	abigen --abi ./gen/contracts/ClaimsManager.json --pkg contracts --type ClaimsManager --out ./gen/contracts/claims_manager.go

	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/eth-contracts/ABIs/DelegateManagerV2.json | jq '.abi' > ./gen/contracts/DelegateManagerV2.json
	abigen --abi ./gen/contracts/DelegateManagerV2.json --pkg contracts --type DelegateManager --out ./gen/contracts/delegate_manager.go

	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/eth-contracts/ABIs/AudiusToken.json | jq '.abi' > ./gen/contracts/AudiusToken.json
	abigen --abi ./gen/contracts/AudiusToken.json --pkg contracts --type AudiusToken --out ./gen/contracts/audius_token.go

	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/eth-contracts/ABIs/EthRewardsManager.json | jq '.abi' > ./gen/contracts/EthRewardsManager.json
	abigen --abi ./gen/contracts/EthRewardsManager.json --pkg contracts --type EthRewardsManager --out ./gen/contracts/eth_rewards_manager.go


build:
	go build

format:
	go fmt ./...

test:
	go test ./... 

hotreload:
	modd

test_verbose:
	go test -v ./...

example:
	go mod tidy
	make build
	go run ./examples/$(EXAMPLE)/main.go
