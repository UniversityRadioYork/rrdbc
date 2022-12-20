package panel

import (
	"net/http"
	"text/template"

	"github.com/UniversityRadioYork/rrdbc/pkg/switcher"
)

type MCRPanel struct {
	SourcesWidth      int
	DestinationsWidth int

	SourcesAndDestinations switcher.SourcesAndDestinations

	SourceGrid      [][]string
	DestinationGrid [][]string
}

type RenderButton struct {
	ID   string
	Name string
}

func (p *MCRPanel) CreateSourceButtonGrid() [][]RenderButton {
	rows := [][]RenderButton{}

	for _, row := range p.SourceGrid {
		buttonRow := []RenderButton{}
		for _, source := range row {
			if source == "" {
				buttonRow = append(buttonRow, RenderButton{"", ""})
			}

			for id, s := range p.SourcesAndDestinations.Sources {
				if s.GetName() == source {
					buttonRow = append(buttonRow, RenderButton{id.String(), s.GetName()})
					break
				}
			}
		}
		rows = append(rows, buttonRow)
	}

	return rows
}

func (p *MCRPanel) CreateDestinationButtonGrid() [][]RenderButton {
	rows := [][]RenderButton{}

	for _, row := range p.DestinationGrid {
		buttonRow := []RenderButton{}
		for _, dest := range row {
			if dest == "" {
				buttonRow = append(buttonRow, RenderButton{"", ""})
				continue
			}

			for id, d := range p.SourcesAndDestinations.Destinations {
				if d.GetName() == dest {
					buttonRow = append(buttonRow, RenderButton{id.String(), d.GetName()})
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
		Sources      [][]RenderButton
		Destinations [][]RenderButton
	}{
		Sources:      p.CreateSourceButtonGrid(),
		Destinations: p.CreateDestinationButtonGrid(),
	}); err != nil {
		// TODO
		panic(err)
	}
}
