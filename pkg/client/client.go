package main

import (
	"context"
	"flag"
	"log"
	"time"

	api "github.com/stevancvetkovic/go-grpc-addressbook/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultFirstName = "Peter"
	defaultLastName  = "Pan"
)

var (
	addr      = flag.String("addr", "127.0.0.1:50051", "the address to connect to")
	firstname = flag.String("firstname", defaultFirstName, "Firstname")
	lastname  = flag.String("lastname", defaultLastName, "Lastname")
)

type Client struct {
	cc  *grpc.ClientConn
	rpc api.AddressbookClient
}

func New(endpoint string, opts ...grpc.DialOption) (c *Client, err error) {
	c = &Client{}

	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c.rpc = api.NewAddressbookClient(c.cc)

	return c, nil
}

func (c *Client) Close() error {
	return c.cc.Close()
}

func (c *Client) GetAddress(firstname, lastname string) (*api.AddressResponse, error) {
	request := &api.AddressRequest{
		FirstName: firstname,
		LastName:  lastname,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.rpc.GetAddress(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func main() {
	flag.Parse()

	client, err := New("127.0.0.1:50051")
	if err != nil {
		panic(err)
	}

	r, err := client.GetAddress(*firstname, *lastname)
	if err != nil {
		panic(err)
	}

	log.Printf("Street: %s", r.GetStreet())
	log.Printf("City: %s", r.GetCity())
}
