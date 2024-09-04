package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/stevancvetkovic/go-grpc-addressbook"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedAddressBookServer
}

func (s *server) GetAddress(_ context.Context, in *pb.AddressRequest) (*pb.AddressResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.AddressResponse{
		Street: "Beogradska",
		City:   "Belgrade",
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAddressbookServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
