package mock

import (
	"context"
	"os"
	"fmt"
	"encoding/json"

	api "github.com/stevancvetkovic/go-grpc-addressbook/pkg/api/v1"
	"github.com/stevancvetkovic/go-grpc-addressbook/pkg/utils"
	"google.golang.org/grpc"
)

const (
	GetAddressRPC = "addressbook.v1.Addressbook/GetAddress"
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
	OnGetAddress func(context.Context, *api.AddressRequest) (*api.AddressResponse, error)
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

	s.OnGetAddress = nil
}

// UseFixture loadsa a JSON fixture from disk (usually in a testdata folder) to use as
// the protocol buffer response to the specified RPC, simplifying handler mocking.
func (s *AddressbookServer) UseFixture(rpc, path string) (err error) {
	var data []byte
	if data, err = os.ReadFile(path); err != nil {
		return fmt.Errorf("could not read fixture: %v", err)
	}

	switch rpc {
	case GetAddressRPC:
		out := &api.AddressResponse{}
		if err = json.Unmarshal(data, out); err != nil {
			return fmt.Errorf("could not unmarshal json into %T: %v", out, err)
		}
		s.OnGetAddress = func(context.Context, *api.AddressRequest) (*api.AddressResponse, error) {
			return out, nil
		}
	default:
		return fmt.Errorf("unknown RPC %q", rpc)
	}

	return nil
}

func (s *AddressbookServer) GetAddress(ctx context.Context, in *api.AddressRequest) (out *api.AddressResponse, err error) {
	s.Calls[GetAddressRPC]++
	return s.OnGetAddress(ctx, in)
}
