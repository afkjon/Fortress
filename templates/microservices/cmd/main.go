package main

import (
	"context"
	"log"
	"net"

	pb "microservices/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TCP Port
const port = ":50051"

type server struct {
	pb.UnimplementedYourServiceServer
}

// Sample gRPC Method
func (s *server) HandleMethod(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
	return &pb.YourResponse{
		Message: "Hello, " + req.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
