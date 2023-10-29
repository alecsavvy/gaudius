package gaudius

type AudiusSdk struct {
	Discovery *DiscoveryNode
	Storage *StorageNode
}

func NewSdk() *AudiusSdk {
	discovery := NewDiscoveryNode("https://discoveryprovider3.audius.co")
	storage := NewStorageNode("https://creatornode.audius.co")
	return &AudiusSdk{ Discovery: discovery, Storage: storage}
}
