package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"LearMicroserviceBasic/microservices/common-services/pb"
	"google.golang.org/grpc"
)

type Common struct {
	pb.UnimplementedCommonServiceServer
}

func (_self Common) GetInformation(ctx context.Context, req *pb.CommonRequest) (*pb.CommonResponse, error) {
	fmt.Printf("receive a Common request with ID: %v", req.GetId())
	//Implement Call microservice
	accountConn, err := grpc.DialContext(ctx, os.Getenv("ACCOUNT_HOST")+":"+os.Getenv("ACCOUNT_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err.Error())
	}
	accountResponse, err := _self.GetAccount(ctx, req.Id, accountConn)
	if err != nil {
		return nil, err
	}

	activityConn, err := grpc.DialContext(ctx, os.Getenv("ACTIVITY_HOST")+":"+os.Getenv("ACTIVITY_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err.Error())
	}
	activityResponse, err := _self.GetActivity(ctx, req.Id, activityConn)
	if err != nil {
		return nil, err
	}

	return &pb.CommonResponse{
		Name:       accountResponse.GetName(),
		Activities: activityResponse.GetActivities(),
	}, nil
}

func (_self Common) GetAccount(ctx context.Context, id int64, conn *grpc.ClientConn) (*pb.AccountResponse, error) {
	accountClient := pb.NewAccountServiceClient(conn)
	res, err := accountClient.GetAccount(ctx, &pb.AccountRequest{
		Id: id,
	})
	log.Print("Send Account ID: ", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (_self Common) GetActivity(ctx context.Context, id int64, conn *grpc.ClientConn) (*pb.ActivityResponse, error) {
	activityClient := pb.NewActivityServiceClient(conn)
	res, err := activityClient.GetActivities(ctx, &pb.ActivityRequest{
		Id: id,
	})
	log.Print("Send Activity ID: ", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
