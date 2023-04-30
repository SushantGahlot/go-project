package database

import (
	"context"
	"errors"
)

type AuthorDAO struct {
	*Database
}

func (a *AuthorDAO) GetAuthorIdsByEmail(ctx context.Context, emails []string) (authorIds []int32, err error) {
	if len(emails) == 0 {
		return nil, errors.New("emails can't be empty")
	}

	q := `
		SELECT authorid FROM author WHERE email IN (?)
	`

	err = a.SelectxInContext(ctx, &authorIds, q, emails)
	if err != nil {
		return nil, err
	}

	return authorIds, nil
}

func (a *AuthorDAO) GetAuthorsByIds(ctx context.Context, authorIds []int32) (authors []*Author, err error) {
	if len(authorIds) == 0 {
		return nil, errors.New("author IDs can not be empty")
	}

	q := `
		SELECT authorid, username, firstname, lastname, email FROM author WHERE authorid IN (?);
	`

	err = a.SelectxInContext(ctx, &authors, q, authorIds)
	if err != nil {
		return nil, err
	}

	return authors, nil
}
