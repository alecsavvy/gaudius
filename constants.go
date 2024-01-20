package gaudius

// contract addresses
var (
	MainnetRegistryAddress = "0xd976d3b4f4e22a238c1A736b6612D22f17b6f64C"
	TestnetRegistryAddress = "0xF27A9c44d7d5DDdA29bC1eeaD94718EeAC1775e3"
)

// contract keys
var (
	RegistryKey               = Utf8ToHex("Registry")
	GovernanceKey             = Utf8ToHex("Governance")
	StakingKey                = Utf8ToHex("StakingProxy")
	ServiceProviderFactoryKey = Utf8ToHex("ServiceProviderFactory")
	ClaimsManagerKey          = Utf8ToHex("ClaimsManagerProxy")
	DelegateManagerKey        = Utf8ToHex("DelegateManager")
	AudiusTokenKey            = Utf8ToHex("Token")
	RewardsManagerKey         = Utf8ToHex("EthRewardsManagerProxy")
)
