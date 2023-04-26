package database

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

type AuthorDAO struct {
	db *Database
}

func (a *AuthorDAO) GetAuthorIdsByEmail(ctx context.Context, emails []string) (authorIds []int32, err error) {
	if len(emails) == 0 {
		return nil, errors.New("emails can't be empty")
	}

	q := `
		SELECT authorid FROM author WHERE id IN (?)
	`

	q, args, err := sqlx.In(q, emails)
	if err != nil {
		return nil, err
	}

	err = a.db.SelectContext(ctx, a.db.DB, &authorIds, q, args...)
	if err != nil {
		return nil, err
	}

	return authorIds, nil
}
