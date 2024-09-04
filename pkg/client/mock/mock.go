package mock

import (
	"context"

	api "github.com/stevancvetkovic/go-grpc-addressbook/api/v1"
	"github.com/stevancvetkovic/go-grpc-addressbook/utils"
	"google.golang.org/grpc"
)

// New creates a new mock RemotePeer. If bufnet is nil, one is created for the user.
func New(bufnet *utils.Listener) *AddressbookServer {
	if bufnet == nil {
		bufnet = utils.New()
	}

	remote := &AddressbookServer{
		bufnet: bufnet,
		srv:    grpc.NewServer(),
		Calls:  make(map[string]int),
	}

	api.RegisterAddressbookServer(remote.srv, remote)
	go remote.srv.Serve(remote.bufnet.Sock())
	return remote
}

type AddressbookServer struct {
	api.UnimplementedAddressbookServer
	bufnet     *utils.Listener
	srv        *grpc.Server
	Calls      map[string]int
}

func (s *AddressbookServer) Channel() *utils.Listener {
	return s.bufnet
}

func (s *AddressbookServer) Shutdown() {
	s.srv.GracefulStop()
	s.bufnet.Close()
}

func (s *AddressbookServer) Reset() {
	for key := range s.Calls {
		s.Calls[key] = 0
	}
}

func (s *AddressbookServer) GetAddress(ctx context.Context, in *api.AddressRequest) (out *api.AddressResponse, err error) {
	return &api.AddressResponse{
		Street: "street",
		City: "city",
	}, nil
}
