package main

import (
	"github.com/UniversityRadioYork/rrdbc/pkg/panel"
	"github.com/UniversityRadioYork/rrdbc/pkg/switcher"
	"github.com/google/uuid"
)

var sources map[uuid.UUID]switcher.Source = map[uuid.UUID]switcher.Source{
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "INT1"}},
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "INT2"}},
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "INT3"}},
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "INT4"}},
	uuid.New(): &switcher.AudioFileSource{BaseSource: switcher.BaseSource{Name: "CONT"}},
	uuid.New(): &switcher.StopSource{BaseSource: switcher.BaseSource{Name: "OFFA"}},
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "EXT1"}},
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "EXT2"}},
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "EXT3"}},
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "EXT4"}},
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "PRES"}},
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "STUD"}},
	uuid.New(): &switcher.AudioStreamSource{BaseSource: switcher.BaseSource{Name: "SUST"}},
	uuid.New(): &switcher.StopSource{BaseSource: switcher.BaseSource{Name: "STOP"}},
}

func addMetaSources(sources map[uuid.UUID]switcher.Source) map[uuid.UUID]switcher.Source {
	for _, grp := range example_meta {
		for _, source := range grp.Members {
			sources[uuid.New()] = &switcher.MetaSource{
				BaseSource: switcher.BaseSource{Name: grp.GrpName + "-" + source.ShortName}, Data: source.LongName}
		}
	}

	return sources
}

func createMetaPages() map[string]map[string][]*panel.RawSource {
	m := map[string]map[string][]*panel.RawSource{
		"META": make(map[string][]*panel.RawSource),
	}

	for _, grp := range example_meta {
		m["META"][grp.GrpName] = []*panel.RawSource{}
		for _, source := range grp.Members {
			m["META"][grp.GrpName] = append(m["META"][grp.GrpName], &panel.RawSource{Name: source.ShortName})
		}
	}

	return m
}

var RosesPanel panel.MCRPanel = panel.MCRPanel{
	SourcesWidth:      4,
	DestinationsWidth: 4,
	Switcher: switcher.Switcher{

		Sources: addMetaSources(sources),

		Destinations: map[uuid.UUID]switcher.Destination{
			uuid.New(): &switcher.StreamDest{BaseDestination: switcher.BaseDestination{Name: "EXT1"}},
			uuid.New(): &switcher.StreamDest{BaseDestination: switcher.BaseDestination{Name: "EXT2"}},
			uuid.New(): &switcher.StreamDest{BaseDestination: switcher.BaseDestination{Name: "EXT3"}},
			uuid.New(): &switcher.StreamDest{BaseDestination: switcher.BaseDestination{Name: "EXT4"}},
			uuid.New(): &switcher.StreamDest{BaseDestination: switcher.BaseDestination{Name: "PRIM"}},
			uuid.New(): &switcher.StreamDest{BaseDestination: switcher.BaseDestination{Name: "SECO"}},
			uuid.New(): &switcher.RecorderDest{BaseDestination: switcher.BaseDestination{Name: "RECA"}},
			uuid.New(): &switcher.RecorderDest{BaseDestination: switcher.BaseDestination{Name: "RECB"}},
			uuid.New(): &switcher.RecorderDest{BaseDestination: switcher.BaseDestination{Name: "RECC"}},
			uuid.New(): &switcher.RecorderDest{BaseDestination: switcher.BaseDestination{Name: "RECD"}},
			uuid.New(): &switcher.RecorderDest{BaseDestination: switcher.BaseDestination{Name: "RECE"}},
			uuid.New(): &switcher.RecorderDest{BaseDestination: switcher.BaseDestination{Name: "RECF"}},
			uuid.New(): &switcher.RecorderDest{BaseDestination: switcher.BaseDestination{Name: "RECG"}},
			uuid.New(): &switcher.RecorderDest{BaseDestination: switcher.BaseDestination{Name: "RECH"}},
			uuid.New(): &switcher.StreamMetaDest{BaseDestination: switcher.BaseDestination{Name: "EXT1 META"}},
			uuid.New(): &switcher.StreamMetaDest{BaseDestination: switcher.BaseDestination{Name: "EXT2 META"}},
			uuid.New(): &switcher.StreamMetaDest{BaseDestination: switcher.BaseDestination{Name: "EXT3 META"}},
			uuid.New(): &switcher.StreamMetaDest{BaseDestination: switcher.BaseDestination{Name: "EXT4 META"}},
			uuid.New(): &switcher.StreamMetaDest{BaseDestination: switcher.BaseDestination{Name: "MAIN META"}},
			uuid.New(): &switcher.RecorderMetaDest{BaseDestination: switcher.BaseDestination{Name: "RECA META"}},
			uuid.New(): &switcher.RecorderMetaDest{BaseDestination: switcher.BaseDestination{Name: "RECB META"}},
			uuid.New(): &switcher.RecorderMetaDest{BaseDestination: switcher.BaseDestination{Name: "RECC META"}},
			uuid.New(): &switcher.RecorderMetaDest{BaseDestination: switcher.BaseDestination{Name: "RECD META"}},
			uuid.New(): &switcher.RecorderMetaDest{BaseDestination: switcher.BaseDestination{Name: "RECE META"}},
			uuid.New(): &switcher.RecorderMetaDest{BaseDestination: switcher.BaseDestination{Name: "RECF META"}},
			uuid.New(): &switcher.RecorderMetaDest{BaseDestination: switcher.BaseDestination{Name: "RECG META"}},
			uuid.New(): &switcher.RecorderMetaDest{BaseDestination: switcher.BaseDestination{Name: "RECH META"}},
		},
	},

	DestinationGrid: [][]string{
		{"EXT1", "EXT2", "EXT3", "EXT4"},
		{"", "", "", ""},
		{"", "", "", ""},
		{"PRIM", "SECO", "", ""},
		{"RECA", "RECB", "RECC", "RECD"},
		{"RECE", "RECF", "RECG", "RECH"},
		{"", "", "", ""},
		{"EXT1 META", "EXT2 META", "EXT3 META", "EXT4 META"},
		{"MAIN META", "", "", ""},
		{"RECA META", "RECB META", "RECC META", "RECD META"},
		{"RECE META", "RECF META", "RECG META", "RECH META"},
	},

	SourceGrid: [][]*panel.RenderButton{
		{
			{
				Name:       "INT1",
				ButtonType: panel.RawButtonType,
			},
			{
				Name:       "INT2",
				ButtonType: panel.RawButtonType,
			}, {
				Name:       "INT3",
				ButtonType: panel.RawButtonType,
			}, {
				Name:       "INT4",
				ButtonType: panel.RawButtonType,
			}},
		{
			{
				Name:       "CONT",
				ButtonType: panel.RawButtonType,
			},
			{}, {}, {
				Name:       "OFFA",
				ButtonType: panel.RawButtonType,
			}}, {{}, {}, {}, {}}, {
			{
				Name:       "EXT1",
				ButtonType: panel.RawButtonType,
			},
			{
				Name:       "EXT2",
				ButtonType: panel.RawButtonType,
			},
			{
				Name:       "EXT3",
				ButtonType: panel.RawButtonType,
			},
			{
				Name:       "EXT4",
				ButtonType: panel.RawButtonType,
			},
		},
		{
			{
				Name:       "PRES",
				ButtonType: panel.RawButtonType,
			}, {
				Name:       "STUD",
				ButtonType: panel.RawButtonType,
			},
			{
				Name:       "SUST",
				ButtonType: panel.RawButtonType,
			}, {},
		}, {
			{}, {}, {}, {
				Name:       "STOP",
				ButtonType: panel.RawButtonType,
			},
		}, {
			{}, {}, {}, {},
		}, {
			{
				Name:       "1",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "2",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "3",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "4",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
		},
		{
			{
				Name:       "5",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "6",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "7",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "8",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
		},
		{
			{
				Name:       "9",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "10",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "11",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "12",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
		},
		{
			{
				Name:       "13",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "14",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "15",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "16",
				ButtonType: panel.PostPageButtonType,
				PageGroup:  "META",
			},
		}, {
			{}, {}, {}, {},
		}, {
			{
				Name:       "GEN",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "HBALL",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "FOOT",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "HOCK",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
		},
		{
			{
				Name:       "LCROS",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "VOLLEY",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "NETB",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "SWIM",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
		},
		{
			{
				Name:       "BADMIN",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "DARTS",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "BASKET",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "CLIMB",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
		},
		{
			{
				Name:       "FRISB",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "RUGBY",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{
				Name:       "OTHER",
				ButtonType: panel.PageButtonType,
				PageGroup:  "META",
			},
			{},
		},
	},

	SourcePages: createMetaPages(),
}
