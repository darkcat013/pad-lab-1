go_protoc:
	protoc --go_out=./Gateway/services --go-grpc_out=./Gateway/services  ./Gateway/proto/owner.proto
	protoc --go_out=./Gateway/services --go-grpc_out=./Gateway/services  ./Gateway/proto/veterinary.proto
	protoc --go_out=./Gateway/services --go-grpc_out=./Gateway/services  ./Gateway/proto/test.proto
	