package main

import (
	pb "first-go-project/api/generated"
	"first-go-project/internal/app/database"
	"first-go-project/internal/app/server"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("count not connect to database %s", err.Error())
	}
	authorDAO := database.AuthorDAO{Database: db}
	postDAO := database.PostDAO{Database: db}
	authorServer := server.AuthorServer{DB: &authorDAO, UnimplementedAuthorServiceServer: pb.UnimplementedAuthorServiceServer{}}
	postServer := server.PostServer{AuthorDAO: &authorDAO, PostDAO: &postDAO, UnimplementedPostServiceServer: pb.UnimplementedPostServiceServer{}}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterAuthorServiceServer(grpcServer, &authorServer)
	pb.RegisterPostServiceServer(grpcServer, &postServer)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 50051))
	if err != nil {
		log.Fatalf("failed to start gRPC server, error: %s", err.Error())
	}

	grpcServer.Serve(lis)
}
