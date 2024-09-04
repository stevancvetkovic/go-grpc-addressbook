package client_test

import (
	"testing"

	client "github.com/stevancvetkovic/go-grpc-addressbook/client"
	"github.com/stevancvetkovic/go-grpc-addressbook/client/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetAddress(t *testing.T) {
	// Setup RemoteAgenda mock server
	srv := mock.New(nil)
	defer srv.Shutdown()

	// Create a client to test that is connected to the mock server.
	dialer := grpc.WithContextDialer(srv.Channel().Dialer)
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())

	addressbook, err := client.New("bufconn", dialer, creds)
	require.NoError(t, err, "could not connect to remote addressbook via bufconn")

	r, err := addressbook.GetAddress("hello", "world")
	require.Equal(t, r.City, "city")
	require.NoError(t, err)
}
