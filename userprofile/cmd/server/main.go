package main

import (
	"fmt"
	"github/sahilrana7582/go-grpc-graphql-microservice/userprofile/config"
	"github/sahilrana7582/go-grpc-graphql-microservice/userprofile/db"
	"github/sahilrana7582/go-grpc-graphql-microservice/userprofile/grpc"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode,
	)

	db.InitDB(&dsn)

	grpc.StartGRPCServer("50052")

}
