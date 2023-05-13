# gRPC Example Go

Follow these setup to run the [quick start][] example:

1.  Generate gRPC code:

    ```console
    $ protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      message/message.proto
    ```

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart
