package grpc

import (
	"context"
	pb "github/sahilrana7582/go-grpc-graphql-microservice/userprofile/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedUserProfileServiceServer
}

func (s *Server) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	log.Printf("GetUserProfile called with ID: %d", req.Id)

	// Dummy hardcoded data
	profile := &pb.UserProfile{
		Id:        req.Id,
		Username:  "john_doe",
		Email:     "john.doe@example.com",
		FullName:  "John Doe",
		Bio:       "Software Developer",
		Interests: []string{"coding", "music", "hiking"},
		Address: &pb.Address{
			Street:     "123 Main St",
			City:       "San Francisco",
			State:      "CA",
			Country:    "USA",
			PostalCode: "94105",
		},
		Phones: []*pb.PhoneNumber{
			{Type: pb.PhoneNumber_MOBILE, Number: "+1-123-456-7890"},
			{Type: pb.PhoneNumber_HOME, Number: "+1-098-765-4321"},
		},
		IsActive:  true,
		CreatedAt: timestamppb.New(time.Now().Add(-365 * 24 * time.Hour)),
		UpdatedAt: timestamppb.New(time.Now()),
	}

	return &pb.GetUserProfileResponse{Profile: profile}, nil
}

func (s *Server) SaveUserProfile(ctx context.Context, req *pb.SaveUserProfileRequest) (*pb.SaveUserProfileResponse, error) {
	log.Printf("SaveUserProfile called with profile: %+v", req.Profile)

	now := timestamppb.Now()
	profile := req.Profile
	if profile.CreatedAt == nil {
		profile.CreatedAt = now
	}
	profile.UpdatedAt = now

	return &pb.SaveUserProfileResponse{Profile: profile}, nil
}

func StartGRPCServer(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserProfileServiceServer(grpcServer, &Server{})

	log.Printf("Starting gRPC UserProfile service on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
