package gaudius

type AudiusSdkParams struct {
	DiscoveryNode        string
	CreatorNode          string
	EntityManagerAddress string
	AcdcGatewayRpc       string
}

/*
*
First will determine the environment.

Second will read in defaults for that environment.

Third will read in a dotenv / env vars for that environment.

Fourth will read in hardcoded values from function input.
*/
func InitSdkParams() (*AudiusSdkParams, error) {
	return nil, nil
}
