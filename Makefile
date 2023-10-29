all: gen build test

gen:
	swagger generate model --skip-validation -f https://discoveryprovider3.audius.co/v1/swagger.json -m models -t ./gen/discovery
	swagger generate model --skip-validation -f https://discoveryprovider3.audius.co/v1/full/swagger.json -m models -t ./gen/discoveryFull

build:
	go build

test:
	go test -v ./... 