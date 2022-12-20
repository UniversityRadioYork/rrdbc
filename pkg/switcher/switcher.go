package switcher

import "github.com/google/uuid"

type SourcesAndDestinations struct {
	Sources      map[uuid.UUID]Source
	Destinations map[uuid.UUID]Destination

	AllowedConnections map[string][]string
}

func (sad *SourcesAndDestinations) MapSourceToDestination(sourceID, destinationID uuid.UUID) error {
	if dest, ok := sad.Destinations[destinationID]; ok {
		if source, ok := sad.Sources[sourceID]; ok {

			var connectionAllowed bool
			for grp, mapGroups := range sad.AllowedConnections {
				if _, ok := source.GetGroups()[grp]; ok {
					for _, destGrp := range mapGroups {
						if _, ok := dest.GetGroups()[destGrp]; ok {
							connectionAllowed = true
						}
					}
				}
			}

			if connectionAllowed {
				return dest.SetSource(source)
			}

			// TODO Return Error (connection not allowed)
		}
		// TODO Error (no source with ID)
	}
	// TODO Error (no dest with ID)

	return nil // temp
}
