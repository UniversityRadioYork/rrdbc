package switcher

type Source interface {
	GetName() string
	GetGroups() map[string]bool
}

type BaseSource struct {
	Name   string
	Groups map[string]bool
}

func (b *BaseSource) GetName() string {
	return b.Name
}

func (b *BaseSource) GetGroups() map[string]bool {
	return b.Groups
}

type MetaSource struct {
	BaseSource
	Data string
}
