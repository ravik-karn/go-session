````
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
brew install protoc-gen-go
mkdir grpc-server
protoc proto/grpc-server.proto --go_out=plugins=grpc:grpc-server
````
