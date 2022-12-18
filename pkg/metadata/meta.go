package metadata

import "github.com/UniversityRadioYork/rrdbc/pkg/switcher"

func GetStreamMetadata() map[string]string {
	meta := make(map[string]string)
	for dest, data := range switcher.Destinations {
		if data.DestType == switcher.MetaStreamDestination {
			meta[dest] = data.Source.Data
		}
	}

	return meta
}
