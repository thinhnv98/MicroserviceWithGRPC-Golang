package controller

import (
	"context"
	"fmt"

	"LearMicroserviceBasic/microservices/activity-services/pb"
)

type Activity struct {
	pb.UnimplementedActivityServiceServer
}

func (_self Activity) GetActivities(ctx context.Context, req *pb.ActivityRequest) (*pb.ActivityResponse, error) {
	fmt.Printf("Receive a Activity request %v", req)
	return &pb.ActivityResponse{
		Activities: []string{"Send money", "Withdrawal"},
	}, nil
}
