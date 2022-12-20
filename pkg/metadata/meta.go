package metadata

import (
	"github.com/UniversityRadioYork/rrdbc/pkg/switcher"
	"github.com/google/uuid"
)

func GetStreamMetadata(destinations map[uuid.UUID]switcher.Destination) map[string]string {
	meta := make(map[string]string)
	for _, dest := range destinations {
		if _, ok := dest.(*switcher.StreamMetaDest); ok {
			meta[dest.GetName()] = dest.GetSource().(*switcher.MetaSource).Data
		}
	}

	return meta
}
