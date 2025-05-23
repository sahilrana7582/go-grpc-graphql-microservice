module github.com/sahilrana27582/go-grpc-graphql-microservice/user-service

go 1.23.3

require (
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.38.0
	google.golang.org/grpc v1.72.1
	google.golang.org/protobuf v1.36.5
)

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/segmentio/kafka-go v0.4.48 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
)

replace github.com/sahilrana27582/go-grpc-graphql-microservice/user-service/grpc/pb => ./grpc/pb
