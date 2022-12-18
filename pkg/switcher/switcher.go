package switcher

func MapSourceToDestination(source Source, destination string) error {
	if dest, ok := Destinations[destination]; ok {
		dest.Source = source
		Destinations[destination] = dest
	} else {
		// TODO Error
	}

	return nil
}
