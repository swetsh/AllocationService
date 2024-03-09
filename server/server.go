package main

import (
	pb "allocation-service/proto/allocation/proto"
	service "allocation-service/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// dbManager := &config.DBManager{}
	// if err := dbManager.DatabaseConnection(); err != nil {
	// 	log.Fatalf("failed to connect to database: %s", err)
	// }

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	serverRegistrar := grpc.NewServer()

	pb.RegisterAllocationServiceServer(serverRegistrar, &service.Server{})

	log.Println("Server started...")

	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
