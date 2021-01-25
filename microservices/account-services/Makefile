gen:
	protoc --go_out=./pb \
		   --go-grpc_out=./pb \
		   --grpc-gateway_out=./pb	\
		   ./proto/*.proto

run:
	go run ./served/main.go
