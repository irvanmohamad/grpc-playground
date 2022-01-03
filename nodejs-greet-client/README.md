

node src/client.js


npm install -g grpc-tools
npm i google-protobuf
npm i grpc
npm i @grpc/proto-loader


grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./nodejs-order-client/services/ --grpc_out=grpc_js:./nodejs-order-client/services/ helloworld.proto


grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./services/ --grpc_out=grpc_js:./services/ greet.proto

- https://blog.logrocket.com/creating-a-crud-api-with-node-express-and-grpc/
- https://github.com/grpc/grpc/blob/master/examples/node/static_codegen/greeter_client.js
- https://github.com/grpc/grpc/tree/master/examples/node/static_codegen
- https://github.com/RedKenrok/node-audiorecorder