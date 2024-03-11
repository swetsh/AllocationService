package service

import (
	"allocation-service/api"
	pb "allocation-service/proto/allocation/proto"
	"allocation-service/responses"
	"context"
	"errors"
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

	var orderResponse *responses.OrderResponse

	for _, deliveryPerson := range deliveryPersons {
		if deliveryPerson.OrderID == -1 {
			orderResponse, err = api.PutOrder("http://localhost:8081/api/v1/delivery-persons", deliveryPerson.ID, int(req.Id))
			fmt.Println(orderResponse, err)
		}
	}

	if orderResponse == nil {
		return nil, errors.New("no available delivery person to assign the order")
	}

	return &pb.OrderResponse{
		Id:               int64(orderResponse.OrderID),
		DeliveryPersonId: int64(orderResponse.ID),
		Status:           pb.OrderStatus_ASSIGNED,
	}, nil
}
