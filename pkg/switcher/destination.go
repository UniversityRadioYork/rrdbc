package switcher

type Destination interface {
	GetName() string
	GetGroups() map[string]bool
	GetSource() Source
	SetSource(s Source) error
}

type BaseDestination struct {
	Name   string
	Groups map[string]bool
	Source Source
}

func (b *BaseDestination) GetName() string {
	return b.Name
}

func (b *BaseDestination) GetGroups() map[string]bool {
	return b.Groups
}

func (b *BaseDestination) GetSource() Source {
	return b.Source
}

type StreamMetaDest struct {
	BaseDestination
}

func (s *StreamMetaDest) SetSource(source Source) error {
	// TODO Source Type Checking
	s.Source = source
	return nil
}

type StreamDest struct {
	BaseDestination
}

func (s *StreamDest) SetSource(source Source) error {
	// TODO Source Type Checking
	s.Source = source
	return nil
}

type RecorderDest struct {
	BaseDestination
}

func (r *RecorderDest) SetSource(source Source) error {
	// TODO Source Type Checking
	r.Source = source
	return nil
}

type RecorderMetaDest struct {
	BaseDestination
}

func (r *RecorderMetaDest) SetSource(source Source) error {
	// TODO Source Type Checking
	r.Source = source
	return nil
}
