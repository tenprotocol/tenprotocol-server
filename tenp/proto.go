package tenp

//go:generate protoc -I./ -I./lib ./tenprotocol-server.proto --gogoslick_out=plugins=grpc:./ --swagger_out=../swagger --grpc-gateway_out=./
