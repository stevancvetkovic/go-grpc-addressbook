package client_test

import (
	"testing"

	client "github.com/stevancvetkovic/go-grpc-addressbook/pkg/client"
	"github.com/stevancvetkovic/go-grpc-addressbook/pkg/client/mock"
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

	// _, err = addressbook.GetAddress("hello", "world")
	// require.Error(t, err, "expected an error returned from the server")
	// require.Equal(t, 1, srv.Calls[mock.GetAddressRPC])

	// Test the case where server returns a response
	srv.UseFixture(mock.GetAddressRPC, "testdata/addressresponse.json")
	r, err := addressbook.GetAddress("hello", "world")
	require.NoError(t, err, "expected no error in happy path")
	require.Equal(t, 2, srv.Calls[mock.GetAddressRPC])
	require.Equal(t, r.GetCity(), "Belgrade")
}
