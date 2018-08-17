package main

import (
	"os"

	"github.com/cfx-cv/suez/pkg/server"
	"github.com/cfx-cv/suez/pkg/suez"
)

var (
	channels = []suez.Channel{
		suez.Channel{
			OriginEndpoint: "/distance",
			DestinationURL: os.Getenv("DIJKSTRA_URI"),

			Method: "GET",
		},
		suez.Channel{
			OriginEndpoint: "/staticmap",
			DestinationURL: os.Getenv("NAMI_URI"),

			Method: "GET",
		},
	}

	envs = make(map[string]string)
)

func init() {
	keys := []string{
		"GCP_API_KEY",
	}

	for _, key := range keys {
		envs[key] = os.Getenv(key)
	}
}

func main() {
	server.NewServer(channels, envs).Start()
}
