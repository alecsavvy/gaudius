package gaudius

// env var keys
var (
	EnvironmentKey = "AUDIUS_ENVIRONMENT"
	RpcKey         = "AUDIUS_ETHEREUM_RPC"
	ApiKeyKey      = "AUDIUS_API_KEY"
	ApiSecretKey   = "AUDIUS_PRIVATE_KEY"
	AppNameKey     = "AUDIUS_APP_NAME"
)

// environments
var (
	Mainnet    = "mainnet"
	Testnet    = "testnet"
	Devnet     = "devnet"
	Prod       = "prod"       // mainnet
	Production = "production" // production
	Stage      = "stage"      // testnet
	Staging    = "staging"    // testnet
	Dev        = "dev"        // devnet
	Local      = "local"      // devnet
)

// contract addresses
var (
	MainnetRegistryAddress = "0xd976d3b4f4e22a238c1A736b6612D22f17b6f64C"
	TestnetRegistryAddress = "0xF27A9c44d7d5DDdA29bC1eeaD94718EeAC1775e3"

	MainnetAcdcAddress = "0x1Cd8a543596D499B9b6E7a6eC15ECd2B7857Fd64"
	TestnetAcdcAddress = "0x1Cd8a543596D499B9b6E7a6eC15ECd2B7857Fd64"
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

// cached nodes to make selection faster on init
var (
	MainnetDiscoveryNodesCached = []string{"https://audius-metadata-1.figment.io", "https://audius-metadata-2.figment.io", "https://audius-discovery-1.altego.net", "https://audius-disco.ams-x01.nl.supercache.org", "https://dn-jpn.audius.metadata.fyi", "https://discoveryprovider3.audius.co", "https://discoveryprovider2.audius.co", "https://discoveryprovider.audius.co", "https://audius-metadata-3.figment.io", "https://audius-metadata-4.figment.io", "https://dn1.monophonic.digital", "https://dn-usa.audius.metadata.fyi", "https://discovery-us-01.audius.openplayer.org", "https://dn2.monophonic.digital", "https://audius-discovery-2.altego.net", "https://dn1.nodeoperator.io", "https://audius-disco.dfw-x02.us.supercache.org", "https://audius-discovery-3.altego.net", "https://dn1.matterlightblooming.xyz", "https://audius-dp.singapore.creatorseed.com", "https://discovery.grassfed.network", "https://audius-discovery-1.cultur3stake.com", "https://audius-discovery-3.cultur3stake.com", "https://audius-discovery-4.cultur3stake.com", "https://audius-discovery-5.cultur3stake.com", "https://audius-discovery-7.cultur3stake.com", "https://audius-discovery-8.cultur3stake.com", "https://audius-discovery-9.cultur3stake.com", "https://audius-discovery-10.cultur3stake.com", "https://discovery-au-02.audius.openplayer.org", "https://disc-lon01.audius.hashbeam.com", "https://audius-dp.amsterdam.creatorseed.com", "https://blockdaemon-audius-discovery-01.bdnodes.net", "https://blockdaemon-audius-discovery-02.bdnodes.net", "https://blockdaemon-audius-discovery-03.bdnodes.net", "https://blockdaemon-audius-discovery-04.bdnodes.net", "https://blockdaemon-audius-discovery-05.bdnodes.net", "https://blockdaemon-audius-discovery-06.bdnodes.net", "https://blockdaemon-audius-discovery-07.bdnodes.net", "https://blockchange-audius-discovery-01.bdnodes.net", "https://blockchange-audius-discovery-02.bdnodes.net", "https://blockchange-audius-discovery-03.bdnodes.net", "https://audius-discovery-11.cultur3stake.com", "https://audius-discovery-12.cultur3stake.com", "https://audius-discovery-13.cultur3stake.com", "https://audius-discovery-14.cultur3stake.com", "https://audius-discovery-16.cultur3stake.com", "https://audius-discovery-18.cultur3stake.com", "https://audius-discovery-17.cultur3stake.com", "https://audius-discovery-15.cultur3stake.com", "https://audius-discovery-6.cultur3stake.com", "https://audius-discovery-2.cultur3stake.com", "https://blockdaemon-audius-discovery-08.bdnodes.net", "https://audius-metadata-5.figment.io", "https://dn1.stuffisup.com", "https://audius-discovery-1.theblueprint.xyz", "https://audius-discovery-2.theblueprint.xyz", "https://audius-discovery-3.theblueprint.xyz", "https://audius-discovery-4.theblueprint.xyz", "https://audius.w3coins.io", "https://audius-nodes.com"}
	MainnetContentNodesCached   = []string{"https://creatornode.audius.co", "https://creatornode2.audius.co", "https://creatornode3.audius.co", "https://content-node.audius.co", "https://audius-content-1.figment.io", "https://creatornode.audius.prod-eks-ap-northeast-1.staked.cloud", "https://audius-content-2.figment.io", "https://audius-content-3.figment.io", "https://audius-content-4.figment.io", "https://audius-content-5.figment.io", "https://creatornode.audius1.prod-eks-ap-northeast-1.staked.cloud", "https://creatornode.audius2.prod-eks-ap-northeast-1.staked.cloud", "https://creatornode.audius3.prod-eks-ap-northeast-1.staked.cloud", "https://audius-content-6.figment.io", "https://audius-content-7.figment.io", "https://audius-content-8.figment.io", "https://usermetadata.audius.co", "https://audius-content-9.figment.io", "https://audius-content-10.figment.io", "https://audius-content-11.figment.io", "https://audius.prod.capturealpha.io", "https://content.grassfed.network", "https://blockdaemon-audius-content-01.bdnodes.net", "https://audius-content-1.cultur3stake.com", "https://audius-content-2.cultur3stake.com", "https://audius-content-3.cultur3stake.com", "https://audius-content-4.cultur3stake.com", "https://audius-content-5.cultur3stake.com", "https://audius-content-6.cultur3stake.com", "https://audius-content-7.cultur3stake.com", "https://blockdaemon-audius-content-02.bdnodes.net", "https://blockdaemon-audius-content-03.bdnodes.net", "https://blockdaemon-audius-content-04.bdnodes.net", "https://blockdaemon-audius-content-05.bdnodes.net", "https://blockdaemon-audius-content-06.bdnodes.net", "https://blockdaemon-audius-content-07.bdnodes.net", "https://blockdaemon-audius-content-08.bdnodes.net", "https://blockdaemon-audius-content-09.bdnodes.net", "https://audius-content-8.cultur3stake.com", "https://blockchange-audius-content-01.bdnodes.net", "https://blockchange-audius-content-02.bdnodes.net", "https://blockchange-audius-content-03.bdnodes.net", "https://audius-content-9.cultur3stake.com", "https://audius-content-10.cultur3stake.com", "https://audius-content-11.cultur3stake.com", "https://audius-content-12.cultur3stake.com", "https://audius-content-13.cultur3stake.com", "https://audius-content-14.cultur3stake.com", "https://audius-content-15.cultur3stake.com", "https://audius-content-16.cultur3stake.com", "https://audius-content-17.cultur3stake.com", "https://audius-content-18.cultur3stake.com", "https://audius-content-12.figment.io", "https://cn0.mainnet.audiusindex.org", "https://cn1.mainnet.audiusindex.org", "https://cn2.mainnet.audiusindex.org", "https://cn3.mainnet.audiusindex.org", "https://audius-content-13.figment.io", "https://audius-content-14.figment.io", "https://cn4.mainnet.audiusindex.org", "https://audius-content-1.jollyworld.xyz", "https://audius-creator-1.theblueprint.xyz", "https://audius-creator-2.theblueprint.xyz", "https://audius-creator-3.theblueprint.xyz", "https://audius-creator-4.theblueprint.xyz", "https://audius-creator-5.theblueprint.xyz", "https://audius-creator-6.theblueprint.xyz", "https://creatornode.audius8.prod-eks-ap-northeast-1.staked.cloud", "https://cn1.stuffisup.com"}
	TestnetDiscoveryNodesCached = []string{"https://discoveryprovider2.staging.audius.co", "https://discoveryprovider3.staging.audius.co", "https://discoveryprovider.staging.audius.co", "https://discoveryprovider5.staging.audius.co"}
	TestnetContentNodesCached   = []string{"https://creatornode5.staging.audius.co", "https://creatornode6.staging.audius.co", "https://creatornode7.staging.audius.co", "https://creatornode8.staging.audius.co", "https://creatornode9.staging.audius.co", "https://creatornode10.staging.audius.co", "https://creatornode11.staging.audius.co", "https://creatornode12.staging.audius.co"}
)
