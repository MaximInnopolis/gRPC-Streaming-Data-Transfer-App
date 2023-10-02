proto:
	protoc -I./main_task/protos   --go_out=./main_task/protos   --go-grpc_out=./main_task/protos \
	--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
 	./main_task/protos/service.proto