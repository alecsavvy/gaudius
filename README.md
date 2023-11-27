# gaudius 🎧🎵🎸🎤

golang client for interacting with [audius](https://github.com/AudiusProject/audius-protocol)

## how to use

### library

```
go get github.com/alecsavvy/gaudius
```

```
package main

import "github.com/alecsavvy/gaudius"

func main() {
	sdk := gaudius.NewSdk()
}
```

### examples

```
git clone https://github.com/alecsavvy/gaudius.git
make example tx-subscriber
```

## completed apis

[discovery query api](https://discoveryprovider3.audius.co/v1/swagger.json)

- [x] developer apps
- [x] playlists
- [x] resolve
- [x] tips
- [x] tracks
- [x] users

[transaction subscriber](https://github.com/AudiusProject/audius-protocol/blob/main/packages/discovery-provider/src/tasks/index_nethermind.py)

This is an event stream of incoming transactions from audius through the ACDC network.

## in progress apis

[discovery query full api](https://discoveryprovider3.audius.co/v1/full/swagger.json)

- [ ] playlists
- [ ] reactions
- [ ] search
- [ ] tips
- [ ] tracks
- [ ] transactions
- [ ] users

[storage public api](https://github.com/AudiusProject/audius-protocol/blob/main/mediorum/server/server.go#L280-L315)

- [ ] uploads
- [ ] blob
- [ ] image
- [ ] system
- [ ] delist

[storage internal api](https://github.com/AudiusProject/audius-protocol/blob/main/mediorum/server/server.go#L334-L356)

- [ ] crud
- [ ] blobs
- [ ] metrics
- [ ] system

[discovery write api](https://docs.audius.org/developers/sdk/)

- [ ] developer apps
- [ ] playlists
- [ ] resolve
- [ ] tips
- [ ] tracks
- [ ] users
