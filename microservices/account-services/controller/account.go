package controller

import (
	"context"
	"fmt"

	"LearMicroserviceBasic/microservices/account-services/pb"
)

type AccountController struct {
	pb.UnimplementedAccountServiceServer
}

func (_self AccountController) GetAccount(ctx context.Context, req *pb.AccountRequest) (*pb.AccountResponse, error) {
	fmt.Printf("Receive a Account request %v", req)
	return &pb.AccountResponse{
		Name: "AccountResponse",
	}, nil
}
