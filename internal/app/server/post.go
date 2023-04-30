package server

import (
	"context"
	pb "first-go-project/api/generated"
	db "first-go-project/internal/app/database"
	"log"
	"math"
	_ "math"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

type PostServer struct {
	AuthorDAO *db.AuthorDAO
	PostDAO   *db.PostDAO
	pb.UnimplementedPostServiceServer
}

func (p *PostServer) GetPostsByIds(ctx context.Context, in *pb.GetPostsByIdsRequest) (*pb.GetPostsByIdsResponse, error) {
	if len(in.GetPostIds()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "post ids can not be empty")
	}

	if len(in.GetPostIds()) > 50 {
		return nil, status.Error(codes.InvalidArgument, "post ids can not more than 50")
	}

	if in.GetSearchTerm() == "" {
		return nil, status.Error(codes.InvalidArgument, "search term can not be empty")
	}

	posts, err := p.PostDAO.GetPostsByIds(ctx, in.GetPostIds())
	if err != nil {
		println(err.Error())
		return nil, status.Error(codes.Internal, "failed getting posts")
	}

	if len(posts) == 0 {
		return nil, status.Error(codes.NotFound, "posts not found")
	}

	var (
		post_count_with_term int32
		proto_posts          = make([]*pb.Post, 0, len(posts))
		tfs                  = make([]int32, 0)
	)

	for _, post := range posts {
		var tf int32

		if post.Body.Valid {
			tf = calculate_tf(post.Body.String, in.GetSearchTerm())
		}

		tfs = append(tfs, tf)
		if tf > 0 {
			post_count_with_term += 1
		}

		proto_post_authors := make([]*pb.Author, 0, len(post.AuthorId))

		for i, authorId := range post.AuthorId {
			proto_post_authors = append(proto_post_authors, &pb.Author{
				AuthorId:  authorId,
				FirstName: post.Firstnames[i],
				LastName:  post.Lastnames[i],
				UserName:  post.Usernames[i],
				Email:     post.Emails[i],
			})
		}

		var (
			b string
			u *timestamppb.Timestamp
		)

		if post.Body.Valid {
			b = post.Body.String
		}

		if post.Updated.Valid {
			u = timestamppb.New(post.Updated.Time)
		}

		proto_posts = append(proto_posts, &pb.Post{
			Authors: proto_post_authors,
			Body:    b,
			Title:   post.Title,
			Created: timestamppb.New(post.Created.Time),
			PostId:  post.PostId.Int32,
			Updated: u,
		})
	}

	var idf float32

	if post_count_with_term > 0 {
		idf = float32(math.Log10(float64(len(posts)) / float64(post_count_with_term)))
	}

	for i, proto_post := range proto_posts {
		if tfs[i] == 0 || idf == 0 {
			continue
		}

		proto_post.Tfidf = float32(float32(tfs[i]) * idf)
	}

	return &pb.GetPostsByIdsResponse{Posts: proto_posts}, nil
}

func calculate_tf(body string, term string) int32 {
	if body == "" {
		return 0
	}

	var term_count int32

	for _, word := range strings.Split(body, " ") {
		if word == term {
			term_count += 1
		}
	}

	return term_count

}

func (p *PostServer) UpsertPost(ctx context.Context, in *pb.UpsertPostRequest) (*pb.UpsertPostResponse, error) {
	if in.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "post title can not be empty")
	}

	if len(in.GetAuthorId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "author ids can not be empty")
	}

	err := p.PostDAO.UpsertPost(ctx, in)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "failed upserting post")
	}

	return &pb.UpsertPostResponse{}, nil
}
