package switcher

type destinationType string

const (
	StreamDestination       destinationType = "STREAM"
	RecorderDestination     destinationType = "RECORD"
	MetaStreamDestination   destinationType = "META STREAM"
	MetaRecorderDestination destinationType = "META RECORD"
)

type Destination struct {
	DestType destinationType
	Source   Source
}

var Destinations = map[string]Destination{
	"INT1": {
		DestType: StreamDestination,
	},
	"INT2": {
		DestType: StreamDestination,
	},
	"INT3": {
		DestType: StreamDestination,
	},
	"INT4": {
		DestType: StreamDestination,
	},
	"PRIM": {
		DestType: StreamDestination,
	},
	"SECO": {
		DestType: StreamDestination,
	},
	"RECA": {
		DestType: RecorderDestination,
	},
	"RECB": {
		DestType: RecorderDestination,
	},
	"RECC": {
		DestType: RecorderDestination,
	},
	"RECD": {
		DestType: RecorderDestination,
	},
	"RECE": {
		DestType: RecorderDestination,
	},
	"RECF": {
		DestType: RecorderDestination,
	},
	"RECG": {
		DestType: RecorderDestination,
	},
	"RECH": {
		DestType: RecorderDestination,
	},
	"EXT1 META": {
		DestType: MetaStreamDestination,
	},
	"EXT2 META": {
		DestType: MetaStreamDestination,
	},
	"EXT3 META": {
		DestType: MetaStreamDestination,
	},
	"EXT4 META": {
		DestType: MetaStreamDestination,
	},
	"MAIN META": {
		DestType: MetaStreamDestination,
	},
	"RECA META": {
		DestType: MetaRecorderDestination,
	},
	"RECB META": {
		DestType: MetaRecorderDestination,
	},
	"RECC META": {
		DestType: MetaRecorderDestination,
	},
	"RECD META": {
		DestType: MetaRecorderDestination,
	},
	"RECE META": {
		DestType: MetaRecorderDestination,
	},
	"RECF META": {
		DestType: MetaRecorderDestination,
	},
	"RECG META": {
		DestType: MetaRecorderDestination,
	},
	"RECH META": {
		DestType: MetaRecorderDestination,
	},
}
