package web

import (
	"fmt"
	"github.com/claudioontheweb/witbot/config"
	"github.com/claudioontheweb/witbot/pkg/witbot"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Config	config.Config
}

func NewServer(config config.Config) *Server {
	server := &Server{
		Config: config,
	}
	return server
}

func (s *Server) Run() error {

	wb, err := witbot.NewWitbot(s.Config)
	if err != nil {
		return err
	}

	tr := mux.NewRouter()

	// API
	routes := tr.PathPrefix("/api").Subrouter()

	routes.Use(CORSMiddleware)

	routes.HandleFunc("", wb.HandleReqs)

	fmt.Println("Server listening on: ", s.Config.Port)
	if err := http.ListenAndServe(s.Config.Addr + ":" + s.Config.Port, tr); err != nil {
		return err
	}
	return nil
}
