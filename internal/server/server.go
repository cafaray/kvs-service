package server

import (
	"log"
	"net/http"
	"time"

	"github.com/cafaray/internal/data"
	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
}

// New inicialize a new server with configuration.
func New(port string) (*Server, error) {
	r := chi.NewRouter()

	ur := &UserRouter{Repository: &data.UserRepository{Data: data.New()}}
	er := &ElementRouter{Repository: &data.ElementRepository{Data: data.New()}}

	r.Mount("/api/v1/users", ur.Routes())
	r.Mount("/api/v1/elements", er.Routes())

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server := Server{server: serv}
	return &server, nil
}

// Close server resources.
func (serv *Server) Close() error {
	// TODO: add resource closure
	return nil
}

func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
