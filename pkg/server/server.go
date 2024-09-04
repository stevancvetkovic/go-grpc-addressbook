package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
	api "github.com/stevancvetkovic/go-grpc-addressbook/pkg/api/v1"
	"google.golang.org/grpc"
)

type Server struct {
	api.UnimplementedAddressbookServer
	srv   *grpc.Server
	echan chan error
}

func New() (*Server, error) {
	s := &Server{
		echan: make(chan error, 1),
	}

	// Create the gRPC Server
	s.srv = grpc.NewServer()
	api.RegisterAddressbookServer(s.srv, s)
	return s, nil
}

func (s *Server) Serve(addr string) (err error) {
	// Catch OS signals for graceful shutdowns
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		s.echan <- s.Shutdown()
	}()

	// Listen for TCP requests on the specified address and port
	var sock net.Listener
	if sock, err = net.Listen("tcp", addr); err != nil {
		return fmt.Errorf("could not listen on %q", addr)
	}

	// Run the server
	go s.Run(sock)
	log.Info().Str("listen", addr).Msg("addressbook server started")

	// Block and wait for an error either from shutdown or grpc.
	if err = <-s.echan; err != nil {
		return err
	}
	return nil
}

func (s *Server) Run(sock net.Listener) {
	defer sock.Close()
	if err := s.srv.Serve(sock); err != nil {
		s.echan <- err
	}
}

func (s *Server) Shutdown() error {
	log.Info().Msg("server shutting down")
	s.srv.GracefulStop()
	log.Debug().Msg("server has shut down gracefully")
	return nil
}

func (s *Server) GetAddress(_ context.Context, in *api.AddressRequest) (*api.AddressResponse, error) {
	log.Printf("Received: %s %s", in.GetFirstName(), in.GetLastName())
	return &api.AddressResponse{
		Street: "Beogradska",
		City:   "Belgrade",
	}, nil
}
