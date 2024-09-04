package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/stevancvetkovic/go-grpc-addressbook/addressbook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultFirstName = "Peter"
	defaultLastName  = "Pan"
)

var (
	addr      = flag.String("addr", "127.0.01:50051", "the address to connect to")
	firstname = flag.String("name", defaultFirstName, "Firstname")
	lastname  = flag.String("lastname", defaultLastName, "Lastname")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetAddress(ctx, &pb.AddressRequest{
		FirstName: *firstname,
		LastName:  *lastname,
	})
	if err != nil {
		log.Fatalf("could not get address: %v", err)
	}
	log.Printf("Street: %s", r.GetStreet())
	log.Printf("City: %s", r.GetCity())
}
