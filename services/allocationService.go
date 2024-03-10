package service

import (
	"allocation-service/api"
	pb "allocation-service/proto/allocation/proto"
	"context"
	"fmt"
)

type Server struct {
	pb.UnimplementedAllocationServiceServer
}

func (s *Server) AssignOrderToDeliveryPerson(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {

	deliveryPersons, err := api.GetDeliveryPersons("http://localhost:8081/api/v1/delivery-persons")

	if err != nil {
		return nil, err
	}

	fmt.Println(deliveryPersons, err)

	for _, deliveryPerson := range deliveryPersons {
		if deliveryPerson.OrderID == -1 {
			api.PutOrder("http://localhost:8081/api/v1/delivery-persons", 1, 1)
		}
		fmt.Println(deliveryPerson)
	}

	return &pb.OrderResponse{}, nil
}
