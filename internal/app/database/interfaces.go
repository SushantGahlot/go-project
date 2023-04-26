package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DB interface {
	MustExecContext(ctx context.Context, query string, args ...interface{}) interface{}
	SelectContext(ctx context.Context, q sqlx.QueryerContext, dest interface{}, query string, args ...interface{}) error
}

type AuthorDB interface {
	GetAuthorIdsByEmail(ctx context.Context, emails []string) (authorIds []int32, err error)
	GetAuthorsByIds(ctx context.Context, authorIds []int32) (authors []*Author, err error)
}

type PostDB interface {
	GetPostsByIds(ctx context.Context, postIds []int32) (posts []*Post, err error)
	UpsertPost(ctx context.Context, post Post) error
}
