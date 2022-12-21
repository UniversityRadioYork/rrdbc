package switcher

// Source represents the 'left hand part' of a connection
// for example, an audio stream for distrbution
type Source interface {
	// GetName returns the source name, displayed in the interface
	GetName() string
	// GetGroups returns the group names defined for the source
	// Groups are used to define valid connections to destinations
	GetGroups() map[string]bool
}

// BaseSource defines what any source should have
type BaseSource struct {
	// Name is typically a short string in capitals to represent
	// the source, i.e. in the interface
	Name   string
	Groups map[string]bool
}

func (b *BaseSource) GetName() string {
	return b.Name
}

func (b *BaseSource) GetGroups() map[string]bool {
	return b.Groups
}

// MetaSource represents a source containig textual information
// for distributing, i.e. as a web-player title, on RDS radiotext or
// to save next to a recording
type MetaSource struct {
	BaseSource
	Data string
}
