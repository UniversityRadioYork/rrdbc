package main

// Based on URY's 2022 Roses coverage
var example_meta = []struct {
	GrpName string
	Members []struct {
		ShortName string
		LongName  string
	}
}{
	{
		GrpName: "GEN",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "BACKSOON",
				LongName:  "We'll be back soon with more coverage.",
			},
		},
	},
	{
		GrpName: "HBALL",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "HBALL",
				LongName:  "Handball",
			},
		},
	},
	{
		GrpName: "FOOT",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "COLLSEL",
				LongName:  "Football (College Select)",
			},
			{
				ShortName: "AMRCN",
				LongName:  "American Football",
			},
			{
				ShortName: "COLLA",
				LongName:  "Football (College A)",
			},
			{
				ShortName: "MEN2ND",
				LongName:  "Football (Men's 2nds)",
			},
			{
				ShortName: "WOMEN1",
				LongName:  "Football (Women's 1sts)",
			},
			{
				ShortName: "MEN1",
				LongName:  "Football (Men's 1sts)",
			},
		},
	},
	{
		GrpName: "HOCK",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "MIXDEV",
				LongName:  "Hockey (Mixed Dev)",
			},
			{
				ShortName: "INDOR",
				LongName:  "Indoor Hockey",
			},
		},
	},
	{
		GrpName: "LCROS",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "MEN1",
				LongName:  "Lacrosse (Men's 1sts)",
			},
			{
				ShortName: "WOMEN1",
				LongName:  "Lacrosse (Women's 1sts)",
			},
		},
	},
	{
		GrpName: "VOLLEY",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "MEN",
				LongName:  "Volleyball (Men's)",
			},
			{
				ShortName: "WOMEN",
				LongName:  "Volleyball (Women's)",
			},
			{
				ShortName: "MIXED",
				LongName:  "Volleyball (Mixed)",
			},
		},
	},
	{
		GrpName: "NETB",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "OPNCMNY",
				LongName:  "Opening Ceremony (Netball)",
			},
			{
				ShortName: "COLLSEL",
				LongName:  "Netball (College Select)",
			},
			{
				ShortName: "3RDS",
				LongName:  "Netball 3rds",
			},
			{
				ShortName: "COLLA",
				LongName:  "Netball (College A)",
			},
			{
				ShortName: "WOMEN2",
				LongName:  "Netball (Women's 2nds)",
			},
		},
	},

	{
		GrpName: "SWIM",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "SWIM",
				LongName:  "Swimming",
			},
			{
				ShortName: "WPOLOMEN1",
				LongName:  "Waterpolo (Men's 1sts)",
			},
			{
				ShortName: "WPOLOWMEN1",
				LongName:  "Waterpolo (Women's 1sts)",
			},
			{
				ShortName: "CPOLOP1",
				LongName:  "Canoe Polo (Open 1sts)",
			},
			{
				ShortName: "CPOLOWM1",
				LongName:  "Canoe Polo (Women's 1sts)",
			},
		},
	},

	{
		GrpName: "BADMIN",
		Members: []struct {
			ShortName string
			LongName  string
		}{{
			ShortName: "WOMEN1",
			LongName:  "Badminton (Women's 1st)",
		}},
	},

	{
		GrpName: "DARTS",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "WOMEN1",
				LongName:  "Darts (Women's 1sts)",
			},
			{
				ShortName: "MEN1",
				LongName:  "Darts (Men's 1sts)",
			},
		},
	},

	{
		GrpName: "BASKET",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "WOMEN1",
				LongName:  "Basketball (Women's 1sts)",
			},
			{
				ShortName: "MEN1",
				LongName:  "Basketball (Men's 1sts)",
			},
		},
	},

	{
		GrpName: "CLIMB",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "LEAD",
				LongName:  "Climbing (Lead)",
			},
			{
				ShortName: "BOULD",
				LongName:  "Climbing (Bouldering)",
			},
		},
	},

	{
		GrpName: "FRISB",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "MEN",
				LongName:  "Ultimate Indoor (Men's)",
			},
			{
				ShortName: "WOMEN",
				LongName:  "Ultimate Indoor (Women's)",
			},
			{
				ShortName: "MIXED",
				LongName:  "Ultimate Indoor (Mixed)",
			},
		},
	},

	{
		GrpName: "RUGBY",
		Members: []struct {
			ShortName string
			LongName  string
		}{
			{
				ShortName: "WOMEN1",
				LongName:  "Rugby Union (Women's 1sts)",
			},
			{
				ShortName: "MEN1",
				LongName:  "Rugby Union (Men's 1sts)",
			},
		},
	},

	{
		GrpName: "OTHER",
		Members: []struct {
			ShortName string
			LongName  string
		}{

			{
				ShortName: "FUTS",
				LongName:  "Futsal (Men's 1st)",
			},

			{
				ShortName: "TTENOP1",
				LongName:  "Table Tennis (Open 1sts)",
			},
			{
				ShortName: "CYCLE",
				LongName:  "Cycling",
			},
		},
	},
}
