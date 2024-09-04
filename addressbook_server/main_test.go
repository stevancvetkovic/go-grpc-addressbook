package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/stevancvetkovic/go-grpc-addressbook/addressbook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterGrpcServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestSayHello(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.NewClient("passthrough://bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewGrpcServiceClient(conn)
	resp, err := client.GetAddress(ctx, &pb.AddressRequest{
		FirstName: "Peter",
		LastName:  "Pan",
	})
	if err != nil {
		t.Fatalf("GetAddress failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}
