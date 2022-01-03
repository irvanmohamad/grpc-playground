# go-order-client && go-greet-server
protoc --go_out=./go-greet-client/services --go_opt=paths=source_relative --go-grpc_out=./go-greet-client/services --go-grpc_opt=paths=source_relative greetpb/*.proto
protoc --go_out=./go-greet-server/services --go_opt=paths=source_relative --go-grpc_out=./go-greet-server/services --go-grpc_opt=paths=source_relative greetpb/*.proto


# nodejs
cd nodejs-greet-client && npm install -g grpc-tools
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./nodejs-greet-client/services/ --grpc_out=grpc_js:./nodejs-greet-client/services/ greetpb/*.proto


# python
conda activate grpc
python -m grpc_tools.protoc --proto_path=. --python_out=./python-greet-client  --grpc_python_out=./python-greet-client greetpb/*.proto

