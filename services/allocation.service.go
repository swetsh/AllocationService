package service

import (
	pb "allocation-service/proto/allocation/proto"
	"context"
)

type Server struct {
	pb.UnimplementedAllocationServiceServer
}

func (s *Server) AllocateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {

	return &pb.Order{}, nil
}
