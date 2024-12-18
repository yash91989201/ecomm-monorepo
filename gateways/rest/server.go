package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/yash91989201/ecomm-monorepo/common/clients"
)

type Server struct {
	ctx         context.Context
	port        int
	ecommClient *clients.InventoryClient
	router      http.Handler
	httpServer  *http.Server
}

func NewServer(ctx context.Context, inventoryServiceUrl string, port int) (*Server, error) {
	ecommClient, err := clients.NewInventoryClient(inventoryServiceUrl)
	if err != nil {
		return nil, err
	}

	handler := NewHandler(ctx, ecommClient)

	router := registerRoutes(handler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return &Server{
		ctx:         ctx,
		ecommClient: ecommClient,
		router:      router,
		port:        port,
		httpServer:  server,
	}, nil
}

func (s *Server) Start() error {
	log.Printf("Rest API gateway started on port: %d", s.port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) CloseServiceClients() {
	s.ecommClient.Close()
}
