package main

import "github.com/UniversityRadioYork/rrdbc/pkg/server"

func main() {
	server := server.Server{
		Port: 3000,
		// MetaGroups: example_meta,
		Users: map[string]struct {
			Password string
			Admin    bool
		}{
			"michael": {
				Password: "testABC",
				Admin:    true,
			},
		},
		Panel: RosesPanel,
	}
	server.Start()
}
