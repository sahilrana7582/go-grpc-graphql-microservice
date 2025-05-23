package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/sahilrana27582/go-grpc-graphql-microservice/user-service/config"
	"github.com/sahilrana27582/go-grpc-graphql-microservice/user-service/db"
	"github.com/sahilrana27582/go-grpc-graphql-microservice/user-service/grpc"
	"github.com/sahilrana27582/go-grpc-graphql-microservice/user-service/kafka"
)

func main() {
	kafka.InitKafka("localhost:9092", "user-events")
	cfg := config.Load()

	dbConn, err := sql.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer dbConn.Close()

	repo := db.NewUserRepository(dbConn)

	grpc.StartGRPCServer(cfg.Port, repo)
}
