package switcher

type destinationType string

const (
	StreamDestination   destinationType = "STREAM"
	RecorderDestination destinationType = "RECORD"
	MetaDestination     destinationType = "META"
)

var destinations = map[string]destinationType{
	"INT1": StreamDestination,
	"INT2": StreamDestination,
	"INT3": StreamDestination,
	"INT4": StreamDestination,
}
