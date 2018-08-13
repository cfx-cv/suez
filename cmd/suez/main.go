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
			DestinationURL: os.Getenv("DIJKSTRA_URL"),
		},
	}
)

func main() {
	key := os.Getenv("GCP_API_KEY")
	server.NewServer(channels, key).Start()
}
