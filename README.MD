# gRPC Example Go

## Setup

[grpc] https://grpc.io/docs/languages/go/quickstart

[protobuf] https://protobuf.dev/

1. Install protocol buffer compiler (version 3+)

   ```console
   $ brew install protobuf
   $ protoc --version
   libprotoc 3.21.12
   ```

2. Install the protocol compiler plugins for Go:

   ```console
   $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. Update your PATH so that the protoc compiler can find the plugins:

   ```console
   $ export PATH="$PATH:$(go env GOPATH)/bin"
   ```

## Run Example

1.  Generate gRPC code:

    ```console
    $ protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      message/message.proto
    ```

2.  Run the server:

    ```console
    $ go run grpc_server/main.go
    ```

3.  Run the client:

    ```console
    $ go run grpc_client/main.go --name=Test
    Reply: Hello Test
    ```
