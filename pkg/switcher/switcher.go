package switcher

import (
	"fmt"

	"github.com/google/uuid"
)

var ErrCantCreateConnection error = fmt.Errorf("can't connect that source to that destination")

// Switcher defines what sources, destinations and allowed routes
// can be worked with together.
type Switcher struct {
	Sources      map[uuid.UUID]Source
	Destinations map[uuid.UUID]Destination

	// AllowedConnections works on GROUPS, defining for each source group
	// a list of destination groups that are valid connections to that source
	AllowedConnections map[string][]string
}

// MapSourceToDestination takes IDs for a source and a destination. If joining that
// source to that destination is valid, it'll call the destinations SetSource() method
// passing the source (though this function could reject it later down the chain)
// If it can't connect, it'll return an error
func (sw *Switcher) MapSourceToDestination(sourceID, destinationID uuid.UUID) error {
	dest, ok := sw.Destinations[destinationID]

	if !ok {
		return fmt.Errorf("no destination with ID %v", destinationID)
	}

	source, ok := sw.Sources[sourceID]

	if !ok {
		return fmt.Errorf("no source with ID %v", sourceID)
	}

	var connectionAllowed bool
	for sourceGrp, allowedDestGroups := range sw.AllowedConnections {
		if _, ok := source.GetGroups()[sourceGrp]; !ok {
			// This sourceGrp isn't a group of this source
			continue
		}

		for _, destGroup := range allowedDestGroups {
			if _, ok := dest.GetGroups()[destGroup]; ok {
				// This destGroup is one of the destinations groups
				connectionAllowed = true
				break
			}
		}
		break
	}

	if connectionAllowed {
		return dest.setSource(source)
	}

	return ErrCantCreateConnection

}
