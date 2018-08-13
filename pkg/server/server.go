package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cfx-cv/suez/pkg/suez"
)

type Server struct {
	channels []suez.Channel
	key      string
}

func NewServer(channels []suez.Channel, key string) *Server {
	return &Server{channels: channels, key: key}
}

func (s *Server) Start() {
	router := mux.NewRouter()
	for _, channel := range s.channels {
		router.HandleFunc(channel.OriginEndpoint, channel.HandlerFunc(s.key))
	}

	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal(err)
	}
}
