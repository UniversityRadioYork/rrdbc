package switcher

type sourceType string

const (
	AudioSourceType sourceType = "AUDIO"
	MetaSourceType  sourceType = "META"
)

type Source struct {
	Type      sourceType `json:"type"`
	ShortName string     `json:"shortName"`
	Data      string     `json:"data"`
}
