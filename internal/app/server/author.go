package server

import (
	"context"
	pb "first-go-project/api/generated"
	db "first-go-project/internal/app/database"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthorServer struct {
	DB *db.AuthorDAO
	pb.UnimplementedAuthorServiceServer
}

func (a *AuthorServer) GetAuthorsByIds(ctx context.Context, in *pb.GetAuthorsByIdsRequest) (*pb.GetAuthorsByIdsResponse, error) {
	if len(in.GetAuthorIds()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "author ids can not be empty")
	}

	authors, err := a.DB.GetAuthorsByIds(ctx, in.GetAuthorIds())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "failed fetching authors")
	}

	if len(authors) == 0 {
		return nil, status.Error(codes.NotFound, "authors not found")
	}

	proto_authors := make([]*pb.Author, 0, len(authors))
	for _, author := range authors {
		proto_authors = append(proto_authors, &pb.Author{
			UserName:  author.Username,
			FirstName: author.Firstname,
			LastName:  author.Lastname,
			Email:     author.Email,
			AuthorId:  author.AuthorId,
		})
	}

	return &pb.GetAuthorsByIdsResponse{Authors: proto_authors}, nil
}

func (a *AuthorServer) GetAuthorIdsByEmails(ctx context.Context, in *pb.GetAuthorIdsByEmailsRequest) (*pb.GetAuthorIdsByEmailResponse, error) {
	if len(in.GetEmail()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "email ids can not be empty")
	}

	authorIds, err := a.DB.GetAuthorIdsByEmail(ctx, in.GetEmail())
	if err != nil {
		println(err)
		return nil, status.Error(codes.Internal, "failed fetching emails")
	}

	if len(authorIds) == 0 {
		return nil, status.Error(codes.NotFound, "authors not found")
	}

	return &pb.GetAuthorIdsByEmailResponse{AuthorIds: authorIds}, nil
}
