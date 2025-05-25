package main

import (
	"context"
	"fmt"
	"github/sahilrana7582/go-grpc-graphql-microservice/userprofile/config"
	"github/sahilrana7582/go-grpc-graphql-microservice/userprofile/db"
	"github/sahilrana7582/go-grpc-graphql-microservice/userprofile/grpc"
	pb "github/sahilrana7582/go-grpc-graphql-microservice/userprofile/proto"
	"log"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_conn "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode,
	)
	db.InitDB(&dsn)

	go grpc.StartGRPCServer("50052")

	time.Sleep(time.Second * 1)

	conn, err := grpc_conn.NewClient(
		"localhost:50052",
		grpc_conn.WithInsecure(),
		grpc_conn.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(5),
				grpc_retry.WithBackoff(grpc_retry.BackoffLinear(2*time.Second)),
				grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			),
		),
	)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserProfileServiceClient(conn)

	req := &pb.GetUserProfileRequest{Id: 1}
	resp, err := client.GetUserProfile(context.Background(), req)
	if err != nil {
		log.Fatalf("GetUserProfile failed: %v", err)
	}

	log.Printf("UserProfile received: %+v", resp.Profile)
}
