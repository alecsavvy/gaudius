package gaudius

type Genres struct {
	AllGenres        string
	Electronic       string
	Rock             string
	Metal            string
	Alternative      string
	HipHopRap        string
	Experimental     string
	Punk             string
	Folk             string
	Pop              string
	Ambient          string
	Soundtrack       string
	World            string
	Jazz             string
	Acoustic         string
	Funk             string
	RandBSoul        string
	Devotional       string
	Classical        string
	Reggae           string
	Podcasts         string
	Country          string
	SpokenWord       string
	Comedy           string
	Blues            string
	Kids             string
	Audiobooks       string
	Latin            string
	LoFi             string
	Hyperpop         string
	Techno           string
	Trap             string
	House            string
	TechHouse        string
	DeepHouse        string
	Disco            string
	Electro          string
	Jungle           string
	ProgressiveHouse string
	Hardstyle        string
	GlitchHop        string
	Trance           string
	FutureBass       string
	FutureHouse      string
	TropicalHouse    string
	Downtempo        string
	DrumAndBass      string
	Dubstep          string
	JerseyClub       string
	Vaporwave        string
	Moombahton       string
}

// Genre constants from the audius-protocol repository
// https://github.com/AudiusProject/audius-protocol/blob/23edd8144acc13da2b8d9000f2245e038b438874/packages/libs/src/sdk/types/Genre.ts#L4
var Genre = Genres{
	AllGenres:        "All Genres",
	Electronic:       "Electronic",
	Rock:             "Rock",
	Metal:            "Metal",
	Alternative:      "Alternative",
	HipHopRap:        "Hip-Hop/Rap",
	Experimental:     "Experimental",
	Punk:             "Punk",
	Folk:             "Folk",
	Pop:              "Pop",
	Ambient:          "Ambient",
	Soundtrack:       "Soundtrack",
	World:            "World",
	Jazz:             "Jazz",
	Acoustic:         "Acoustic",
	Funk:             "Funk",
	RandBSoul:        "R&B/Soul",
	Devotional:       "Devotional",
	Classical:        "Classical",
	Reggae:           "Reggae",
	Podcasts:         "Podcasts",
	Country:          "Country",
	SpokenWord:       "Spoken Word",
	Comedy:           "Comedy",
	Blues:            "Blues",
	Kids:             "Kids",
	Audiobooks:       "Audiobooks",
	Latin:            "Latin",
	LoFi:             "Lo-Fi",
	Hyperpop:         "Hyperpop",
	Techno:           "Techno",
	Trap:             "Trap",
	House:            "House",
	TechHouse:        "Tech House",
	DeepHouse:        "Deep House",
	Disco:            "Disco",
	Electro:          "Electro",
	Jungle:           "Jungle",
	ProgressiveHouse: "Progressive House",
	Hardstyle:        "Hardstyle",
	GlitchHop:        "Glitch Hop",
	Trance:           "Trance",
	FutureBass:       "Future Bass",
	FutureHouse:      "Future House",
	TropicalHouse:    "Tropical House",
	Downtempo:        "Downtempo",
	DrumAndBass:      "Drum & Bass",
	Dubstep:          "Dubstep",
	JerseyClub:       "Jersey Club",
	Vaporwave:        "Vaporwave",
	Moombahton:       "Moombahton",
}
