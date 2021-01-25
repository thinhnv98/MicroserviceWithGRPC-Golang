package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"LearMicroserviceBasic/microservices/proxy-server/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	grpcMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterAccountServiceHandlerFromEndpoint(ctx, grpcMux, *flag.String("account-grpc-endpoint", ":"+os.Getenv("ACCOUNT_PORT"), "gRPC server endpoint"), opts); err != nil {
		return err
	}
	if err := pb.RegisterActivityServiceHandlerFromEndpoint(ctx, grpcMux, *flag.String("activity-grpc-endpoint", ":"+os.Getenv("ACTIVITY_PORT"), "gRPC server endpoint"), opts); err != nil {
		return err
	}
	if err := pb.RegisterCommonServiceHandlerFromEndpoint(ctx, grpcMux, *flag.String("common-grpc-endpoint", ":"+os.Getenv("COMMON_PORT"), "gRPC server endpoint"), opts); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Printf("Rest proxy-server serve on \":%s\"", os.Getenv("PORT"))
	return http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		panic(err)
	}
}
