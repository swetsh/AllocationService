package service

import (
	"allocation-service/api"
	pb "allocation-service/proto/allocation/proto"
	"bytes"
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	pb.UnimplementedAllocationServiceServer
}

func PutOrder(orderId int) {
	url := "http://localhost:8081/api/v1/delivery-persons/6"

	// Convert integer to byte slice
	integerBytes := []byte(fmt.Sprintf("%d", orderId))

	// Create a new request with the integer as the request body
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(integerBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status code
	fmt.Println("Response Status:", resp.Status)
}

func (s *Server) AssignOrderToDeliveryPerson(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {

	deliveryPersons, err := api.GetDeliveryPersons("http://localhost:8081/api/v1/delivery-persons")

	if err != nil {
		return nil, err
	}

	fmt.Println(deliveryPersons, err)

	for _, deliveryPerson := range deliveryPersons {
		if deliveryPerson.OrderID == -1 {
			PutOrder(2)
		}
		fmt.Println(deliveryPerson)
	}

	return &pb.OrderResponse{}, nil
}
