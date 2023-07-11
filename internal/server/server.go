package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/khanfromasia/densys/admin/internal/config"
)

type httpServer struct {
	srv *http.Server
	hnd http.Handler
}

func NewHTTPServer(handler http.Handler) *httpServer {
	return &httpServer{
		hnd: handler,
	}
}

func (s *httpServer) Run(cfg config.Config) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Rest.Host, cfg.Rest.Port))

	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	s.srv = &http.Server{
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 40 << 20, // 1 MB
		Handler:        s.hnd,
	}

	log.Println("HTTP server is running on ", fmt.Sprintf("%s:%d", cfg.Rest.Host, cfg.Rest.Port))
	if err = s.srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}

func (s *httpServer) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
