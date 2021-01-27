gen:
	make -C ./grpc gen

run-account-services:
	make -C ../microservices/account-services run

run-activity-services:
	make -C ../microservices/activity-services run

run-common-services:
	make -C ../microservices/common-services run

run-self:
	go run main.go

run-ms:
	make -j 4 run-account-services run-activity-services run-common-services run-self


