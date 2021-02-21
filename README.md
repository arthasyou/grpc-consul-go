# go_grpc_with_consul

### you must install protoc, protoc-gen-go and protoc-gen-go-grpc
### for MacOS run following cmd to install tools

```console
$ brew install protobuf
$ go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
$ export PATH="$PATH:$(go env GOPATH)/bin"
 ```

 ## Usage please watch the test directory