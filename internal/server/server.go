package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler *mux.Router) error {
	s.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		MaxHeaderBytes:    1 << 20,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	log.Printf("Server starting on http://localhost%s\nswagger - http://localhost%s/swagger/index.html", s.httpServer.Addr, s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
