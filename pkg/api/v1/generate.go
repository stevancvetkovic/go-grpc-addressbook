package proto

//go:generate protoc -I=$GOPATH/src/github.com/stevancvetkovic/go-grpc-addressbook/proto --go_out=. --go_opt=module=github.com/stevancvetkovic/go-grpc-addressbook/api/v1 --go-grpc_out=. --go-grpc_opt=module=github.com/stevancvetkovic/go-grpc-addressbook/api/v1 v1/addressbook.proto
