package gaudius

type AudiusSdkParams struct {
	DiscoveryNode        string
	CreatorNode          string
	EntityManagerAddress string
	AcdcGatewayRpc       string
}

func createSdkParams() (*AudiusSdkParams, error) {
	params := &AudiusSdkParams{}

	return params, nil
}
