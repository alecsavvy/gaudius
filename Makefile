all: gen build test

gen:
	curl -s https://discoveryprovider3.audius.co/v1/swagger.json | jq . > ./swagger/discovery.json
	curl -s https://discoveryprovider3.audius.co/v1/full/swagger.json | jq . > ./swagger/discoveryFull.json

build:
	go build

test:
	go test ./...