```
protoc --version                                  
libprotoc 3.11.2
```

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    addressbook.proto
```
