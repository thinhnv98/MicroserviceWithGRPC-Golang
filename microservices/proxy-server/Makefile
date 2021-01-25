gen:
	protoc --go_out=./pb \
		   --go-grpc_out=./pb \
		   --grpc-gateway_out=./pb	\
		   --swagger_out=./swagger \
		   ./proto/*.proto

run-account-services:
	make -C ../account-services run

run-activity-services:
	make -C ../activity-services run

run-common-services:
	make -C ../common-services run

run-self:
	go run main.go

run-ms:
	make -j 4 run-account-services run-activity-services run-common-services run-self
