package panel

import (
	"net/http"
	"text/template"

	"github.com/UniversityRadioYork/rrdbc/pkg/switcher"
	"github.com/google/uuid"
)

type ButtonType string

const (
	RawButtonType      ButtonType = "RAW"
	PostPageButtonType ButtonType = "POSTPAGE"
	PageButtonType     ButtonType = "PAGE"
)

type RenderButton struct {
	ID string

	Name       string
	ButtonType ButtonType

	PageGroup string
}

type RawSource struct {
	Name string
	ID   string
}
type MCRPanel struct {
	SourcesWidth      int
	DestinationsWidth int

	Switcher switcher.Switcher

	SourcePages      map[string]map[string][]*RawSource
	DestinationPages map[string]map[string][]string

	SourceGrid      [][]*RenderButton
	DestinationGrid [][]string
}

func (p *MCRPanel) Populate() {
	// Source Pages
	for _, pages := range p.SourcePages {
		for pageName, sources := range pages {
			for _, source := range sources {
				for id, s := range p.Switcher.Sources {
					if s.GetName() == pageName+"-"+source.Name {
						source.ID = id.String()
					}
				}
			}
		}
	}

	// Source Button Grid Layout
	for _, row := range p.SourceGrid {
		for _, source := range row {
			if source.ID != "" {
				continue
			}

			switch source.ButtonType {
			case RawButtonType:
				for id, s := range p.Switcher.Sources {
					if s.GetName() == source.Name {
						source.ID = id.String()
					}
				}

			case PageButtonType:
				source.ID = uuid.NewString()

			case PostPageButtonType:
				source.ID = source.PageGroup + "-" + source.Name
			}
		}
	}
}

func (p *MCRPanel) createDestinationButtonGrid() [][]RenderButton {
	rows := [][]RenderButton{}

	for _, row := range p.DestinationGrid {
		buttonRow := []RenderButton{}
		for _, dest := range row {
			if dest == "" {
				buttonRow = append(buttonRow, RenderButton{"", "", RawButtonType, ""})
				continue
			}

			for id, d := range p.Switcher.Destinations {
				if d.GetName() == dest {
					buttonRow = append(buttonRow, RenderButton{id.String(), d.GetName(), RawButtonType, ""})
					break
				}
			}
		}
		rows = append(rows, buttonRow)
	}

	return rows
}

func (p *MCRPanel) RenderTemplate(w http.ResponseWriter, tmpltFile string) {
	tmplt, err := template.ParseFiles(tmpltFile)

	if err != nil {
		// TODO
		panic(err)
	}
	if err := tmplt.Execute(w, struct {
		Sources      [][]*RenderButton
		Destinations [][]RenderButton
	}{
		Sources:      p.SourceGrid,
		Destinations: p.createDestinationButtonGrid(),
	}); err != nil {
		// TODO
		panic(err)
	}
}
