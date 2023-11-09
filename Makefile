all: build test

deps:
	brew tap ethereum/ethereum
	brew install ethereum

generate:
	swagger generate model --skip-validation -f https://discoveryprovider3.audius.co/v1/swagger.json -m models -t ./gen/discovery
	swagger generate model --skip-validation -f https://discoveryprovider3.audius.co/v1/full/swagger.json -m models -t ./gen/discoveryFull
	go mod tidy
	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/data-contracts/ABIs/EntityManager.json | jq '.abi' > ./gen/contracts/EntityManager.json
	abigen --abi ./gen/contracts/EntityManager.json --pkg main --type EntityManager --out ./gen/contracts/entity_manager.go

build:
	go build

test:
	go test ./... 

test_verbose:
	go test -v ./...
