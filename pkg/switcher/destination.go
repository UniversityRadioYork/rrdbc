package switcher

// Destination represents the 'right hand side' of a connection, i.e. where
// a source can be sent to. A destination can only have one source linked
// to it at a time.
type Destination interface {
	// GetName returns the destination name, displayed in the interface
	GetName() string

	// GetGroups returns the group names defined for the destination
	// Groups are used to define valid connections from sources
	GetGroups() map[string]bool

	// GetSource will return the currently connected source
	GetSource() Source

	// setSource will attempt to connect the given Source to this
	// destination. It may not be valid (in terms of types rather
	// than groups), so it may return an error
	setSource(s Source) error
}

type BaseDestination struct {
	// Name is typically a short string in capitals to represent
	// the destination, i.e. in the interface
	Name   string
	Groups map[string]bool
	source Source
}

func (b *BaseDestination) GetName() string {
	return b.Name
}

func (b *BaseDestination) GetGroups() map[string]bool {
	return b.Groups
}

func (b *BaseDestination) GetSource() Source {
	return b.source
}

// StreamMetaDest represents a 'live' placeholder for metadata
// provided from a metadata source. These 'live' metadata
// dsetinations are available to be picked up by the metadata API
type StreamMetaDest struct {
	BaseDestination
}

func (s *StreamMetaDest) setSource(source Source) error {
	if _, ok := source.(*MetaSource); !ok {
		return ErrCantCreateConnection
	}

	s.source = source
	return nil
}

// StreamDest TODO comment
type StreamDest struct {
	BaseDestination
}

func (s *StreamDest) setSource(source Source) error {
	// TODO Source Type Checking
	s.source = source
	return nil
}

// RecorderDest TODO comment
type RecorderDest struct {
	BaseDestination
}

func (r *RecorderDest) setSource(source Source) error {
	// TODO Source Type Checking
	r.source = source
	return nil
}

// RecorderMetaDest TODO comment
type RecorderMetaDest struct {
	BaseDestination
}

func (r *RecorderMetaDest) setSource(source Source) error {
	if _, ok := source.(*MetaSource); !ok {
		return ErrCantCreateConnection
	}

	r.source = source
	return nil
}
