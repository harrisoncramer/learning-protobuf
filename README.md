# learning-protobuf

Protocol buffers is an alternative to JSON or XML. It’s like JSON, except it's smaller and faster, and it generates native language bindings.

They are a way of serializing structured data. The definition language is created in `.proto` files. 

The compiler ingests the `.proto` file and autogenerates source code that lets us encode data into protocol buffers.

The compiler can be installed with Brew:

```
$ brew install protobuf
```

You can then run the `protoc` binary to compile our `.proto` files into generated source code.

```
$ protoc --go_out=. --go_opt=paths=source_relative model/book.proto
```

This will output a generated go file into `model/book.pb.go`. Every time you re-run the compiler the generated code will be overwritten.

GRPC will use this protocol: encode data into a byte stream using proto files and streaming it across the network, and then decoding it again with the same schema.
