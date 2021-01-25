package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"LearMicroserviceBasic/microservices/account-services/controller"
	"LearMicroserviceBasic/microservices/account-services/pb"
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

	accountController := controller.AccountController{}

	pb.RegisterAccountServiceServer(grpcServer, accountController)

	go func() {
		listener, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Print("Account service serve at :", os.Getenv("PORT"))
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("failed to start ACCOUNT server", err)
		}
	}()

	// let us wait for an input here (ctrl+c) to stop the client
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v", signal.String())
}
