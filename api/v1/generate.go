package proto

//go:generate protoc -I ../../proto/v1 --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ../../proto/v1/addressbook.proto
