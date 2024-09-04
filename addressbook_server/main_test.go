package main

import (
	"context"
	"log"
	"testing"

	pb "github.com/stevancvetkovic/go-grpc-addressbook/addressbook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetAddress(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.NewClient("passthrough://bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewAddressbookClient(conn)
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
