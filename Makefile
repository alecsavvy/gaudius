EXAMPLE?=get-user-and-image

all: format build test

deps:
	brew tap ethereum/ethereum
	brew install ethereum

generate:
	swagger generate model --skip-validation -f https://discoveryprovider3.audius.co/v1/swagger.json -m models -t ./gen/discovery
	swagger generate model --skip-validation -f https://discoveryprovider3.audius.co/v1/full/swagger.json -m models -t ./gen/discoveryFull
	go mod tidy
	curl -s https://raw.githubusercontent.com/AudiusProject/audius-protocol/main/packages/libs/src/data-contracts/ABIs/EntityManager.json | jq '.abi' > ./gen/contracts/EntityManager.json
	abigen --abi ./gen/contracts/EntityManager.json --pkg contracts --type EntityManager --out ./gen/contracts/entity_manager.go

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
