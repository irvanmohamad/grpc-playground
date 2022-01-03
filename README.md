### Run Go gRPC Server & Client

```
cd go-greet-client && go run src/github.com/mesadhan/client.go
cd go-greet-server && go run src/github.com/mesadhan/server.go
```


# nodejs-greet-client

```
cd nodejs-greet-client && npm start
```


# python-greet-client

```
conda activate grpc
cd python-greet-client && python client.py
```



## Setup .proto to code generate cli

    brew install protoc
    brew install grpc_tools_node_protoc
    brew install protoc-gen-grpc-python
    brew install protoc-gen-grpc-web
