package panel

import (
	"net/http"
	"text/template"

	"github.com/UniversityRadioYork/rrdbc/pkg/switcher"
)

type MCRPanel struct {
	SourcesWidth      int
	DestinationsWidth int

	Switcher switcher.Switcher

	SourceGrid      [][]string
	DestinationGrid [][]string
}

type RenderButton struct {
	ID   string
	Name string
}

func (p *MCRPanel) createSourceButtonGrid() [][]RenderButton {
	rows := [][]RenderButton{}

	for _, row := range p.SourceGrid {
		buttonRow := []RenderButton{}
		for _, source := range row {
			if source == "" {
				buttonRow = append(buttonRow, RenderButton{"", ""})
			}

			for id, s := range p.Switcher.Sources {
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

func (p *MCRPanel) createDestinationButtonGrid() [][]RenderButton {
	rows := [][]RenderButton{}

	for _, row := range p.DestinationGrid {
		buttonRow := []RenderButton{}
		for _, dest := range row {
			if dest == "" {
				buttonRow = append(buttonRow, RenderButton{"", ""})
				continue
			}

			for id, d := range p.Switcher.Destinations {
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
		Sources:      p.createSourceButtonGrid(),
		Destinations: p.createDestinationButtonGrid(),
	}); err != nil {
		// TODO
		panic(err)
	}
}
