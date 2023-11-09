# gaudius

golang client for interacting with audius nodes

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
make example get-user-and-image
```

## apis

[discovery query api](https://discoveryprovider3.audius.co/v1/swagger.json)

- [ ] developer apps
- [ ] playlists
- [ ] resolve
- [ ] tips
- [ ] tracks
- [ ] users

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

[sdk and libs](https://docs.audius.org/developers/sdk/)

- [ ] [login](https://github.com/AudiusProject/audius-protocol/blob/38ac1b7d5d0f87a47b4bf8839f50a9f61ce9845b/packages/libs/src/api/Account.ts#L59)
- [ ] [user sign up](https://github.com/AudiusProject/audius-protocol/blob/38ac1b7d5d0f87a47b4bf8839f50a9f61ce9845b/packages/libs/src/api/Account.ts#L119)
- [ ] [oauth](https://docs.audius.org/developers/sdk-oauth-methods)
- [ ] [hedgehog](https://docs.audius.org/developers/hedgehog)

[discovery write api](https://docs.audius.org/developers/sdk/)

- [ ] developer apps
- [ ] playlists
- [ ] resolve
- [ ] tips
- [ ] tracks
- [ ] users
