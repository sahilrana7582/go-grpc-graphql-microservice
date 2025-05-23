package grpc

import (
	"context"
	"log"

	"net"

	"github.com/google/uuid"
	"github.com/sahilrana27582/go-grpc-graphql-microservice/user-service/db"
	"github.com/sahilrana27582/go-grpc-graphql-microservice/user-service/grpc/pb/github.com/sahilrana27582/go-grpc-graphql-microservice/user-service/grpc/pb"
	"github.com/sahilrana27582/go-grpc-graphql-microservice/user-service/kafka"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *db.UserRepository
}

func NewUserServer(repo *db.UserRepository) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := db.User{
		ID:       uuid.New(),
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	err = kafka.PublishMessage("UserCreated", map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"password": user.Password,
	})

	if err != nil {
		log.Printf("Kafka publish error: %v", err)
		return nil, err
	}

	return &pb.CreateUserResponse{Id: user.ID.String()}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.repo.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:       user.ID.String(),
		Username: user.Username,
		Role:     user.Role,
	}, nil
}

func StartGRPCServer(port string, repo *db.UserRepository) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, NewUserServer(repo))

	log.Printf("Starting gRPC User service on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
