package http

import (
	"context"
	"fmt"

	netHttp "net/http"

	"github.com/gorilla/mux"
	"github.com/kilia/http-example/internal/settings"
)

type Server interface {
	Start(chan<- error)
	Stop(context.Context)error
}

type HttpServer struct {
	engine *netHttp.Server
}

func (s *HttpServer) Start(chErr chan<- error) {
	defer func() {
		if r := recover(); r != nil {
			chErr <- fmt.Errorf("Server error:%v", r)
		}
	}()

	fmt.Printf("Server started:%s\n", s.engine.Addr)
	chErr <- s.engine.ListenAndServe()
}

func (s *HttpServer) Stop(ctx context.Context) error {
	if err := s.engine.Shutdown(ctx); err != nil {
		fmt.Println("Server stopped:", err)
		return err
	}

	fmt.Println("Server stopped gracefully")
	return nil
}	

func NewServer(s *settings.Settings) *HttpServer {
	r := mux.NewRouter()
	r.HandleFunc("/ping", SystemHandler)
	return &HttpServer{
		engine: &netHttp.Server{
			Addr: fmt.Sprintf("%v:%d", s.Server.Host, s.Server.Port),
			Handler: r,
		},
	}
}