package database

import (
	"context"
	pb "first-go-project/api/generated"
)

type DB interface {
	SelectxInContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExecxContext(ctx context.Context, query string, arg interface{}) (interface{}, error)
}

type AuthorDB interface {
	GetAuthorIdsByEmail(ctx context.Context, emails []string) (authorIds []int32, err error)
	GetAuthorsByIds(ctx context.Context, authorIds []int32) (authors []*Author, err error)
}

type PostDB interface {
	GetPostsByIds(ctx context.Context, postIds []int32) (posts []*Post, err error)
	UpsertPost(ctx context.Context, post *pb.UpsertPostRequest) error
}
