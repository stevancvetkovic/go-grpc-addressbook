package mock

import (
	pb "github.com/stevancvetkovic/go-grpc-addressbook/api/v1"
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

	pb.RegisterAddressbookServer(remote.srv, remote)
	go remote.srv.Serve(remote.bufnet.Sock())
	return remote
}

type AddressbookServer struct {
	pb.UnimplementedAddressbookServer
	bufnet *utils.Listener
	srv    *grpc.Server
	Calls  map[string]int
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
