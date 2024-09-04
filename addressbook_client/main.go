/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
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
