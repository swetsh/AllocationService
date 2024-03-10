package service

import (
	pb "allocation-service/proto/allocation/proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	pb.UnimplementedAllocationServiceServer
}

type Location struct {
	GeoCoordinate string `json:"geoCoordinate"`
}

type DeliveryPerson struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
	OrderID  int      `json:"orderId"`
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

	resp, err := http.Get("http://localhost:8081/api/v1/delivery-persons")

	if err != nil {
		fmt.Printf("Error %s\n", err)
		return nil, err
	}

	var deliveryPersons []*DeliveryPerson
	if err := json.NewDecoder(resp.Body).Decode(&deliveryPersons); err != nil {
		return nil, err
	}

	for _, deliveryPerson := range deliveryPersons {
		if deliveryPerson.OrderID == -1 {
			PutOrder(2)
		}
		fmt.Println(deliveryPerson)
	}

	return &pb.OrderResponse{}, nil
}
