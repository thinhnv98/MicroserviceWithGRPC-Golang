package main

import (
	"context"
	"flag"
	"log"
	"os"

	"LearMicroserviceBasic/microservices/proxy-server/reverse-proxy"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func run() error {
	// Instance Server
	server := gin.Default()

	//gRPC Gate-way
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	reverseProxy := reverse_proxy.ReverseProxy{}
	gateway, err := reverseProxy.RegisterProxy(ctx)
	if err != nil {
		panic(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	server.Any("/*proxyPath", gin.WrapH(gateway))
	return server.Run(":" + os.Getenv("PORT"))
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	flag.Parse()
	//
	if err := run(); err != nil {
		panic(err)
	}
}
