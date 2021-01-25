package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"LearMicroserviceBasic/microservices/common-services/controller"
	"LearMicroserviceBasic/microservices/common-services/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	commonController := controller.Common{}

	pb.RegisterCommonServiceServer(grpcServer, commonController)

	go func() {
		listener, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Print("Common service serve at :", os.Getenv("PORT"))
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("failed to start COMMON server", err)
		}
	}()

	// let us wait for an input here (ctrl+c) to stop the client
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v", signal.String())
}
