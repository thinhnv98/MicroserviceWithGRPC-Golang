package reverse_proxy

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"net/http"
	"os"

	"LearMicroserviceBasic/microservices/proxy-server/grpc/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type ReverseProxy struct {
}

func (_self ReverseProxy) RegisterProxy(ctx context.Context) (http.Handler, error) {
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	grpcMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterCommonServiceHandlerFromEndpoint(ctx, grpcMux, *flag.String("common-grpc-endpoint", ":"+os.Getenv("COMMON_PORT"), "gRPC server endpoint"), opts); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	return mux, nil
}
