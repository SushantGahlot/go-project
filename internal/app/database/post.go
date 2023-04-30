package database

import (
	"context"
	"database/sql"
	"errors"
	"log"

	pb "first-go-project/api/generated"

	"github.com/jmoiron/sqlx"
)

type PostDAO struct {
	*Database
}

func (p *PostDAO) GetPostsByIds(ctx context.Context, postIds []int32) ([]PostById, error) {
	if len(postIds) == 0 {
		return nil, errors.New("post IDs can not be empty")
	}

	q := `
		SELECT
			post.postid,
			post.title,
			post.body,
			post.updated,
			post.created,
			array_agg(author.authorid) as authorids,
			array_agg(author.firstname) as firstnames,
			array_agg(author.lastname) as lastnames,
			array_agg(author.username) as usernames,
			array_agg(author.email) as emails
		FROM
			author_post
			INNER JOIN post ON author_post.postid = post.postid
			INNER JOIN author ON author_post.authorid = author.authorid
		WHERE
			post.postid IN (?)
		GROUP BY
			post.postid,
			post.body,
			post.created,
			post.title,
			post.updated
	`
	q, args, err := sqlx.In(q, postIds)
	if err != nil {
		return nil, err
	}

	dbPosts := make([]PostById, 0, len(postIds))

	err = p.SelectxInContext(ctx, &dbPosts, q, args...)
	if err != nil {
		return nil, err
	}

	return dbPosts, nil
}

func (p *PostDAO) UpsertPost(ctx context.Context, post *pb.UpsertPostRequest) error {
	if post.GetTitle() == "" {
		return errors.New("post title can not be empty")
	}

	if len(post.GetAuthorId()) == 0 {
		return errors.New("author IDs can not be empty")
	}

	dbPost := Post{
		Body:   sql.NullString{String: post.GetBody(), Valid: post.GetBody() != ""},
		PostId: sql.NullInt32{Int32: post.GetPostId(), Valid: post.GetPostId() != 0},
		Title:  post.GetTitle(),
	}

	trx, err := p.GetTxContext(ctx)
	if err != nil {
		log.Println("failed getting transaction for upsert")
		return errors.New("failed getting transaction")
	}

	q := `
		INSERT INTO post (body, created, postid, title)
		VALUES (:body, now(), :postid, :title)
		ON CONFLICT(postid)
		DO UPDATE SET body = EXCLUDED.body, title = EXCLUDED.title, updated = now()
		returning postid
	`
	rows, err := trx.NamedQuery(q, dbPost)
	if err != nil {
		trx.Rollback()
		return err
	}

	var postid int32

	for rows.Next() {
		if err = rows.Scan(&postid); err != nil {
			trx.Rollback()
			return err
		}
	}

	author_post := make([]AuthorPost, 0, len(post.GetAuthorId()))

	for _, aid := range post.GetAuthorId() {
		author_post = append(author_post, AuthorPost{sql.NullInt32{}, aid, int32(postid)})
	}

	q = `DELETE FROM author_post WHERE postid = $1`
	_, err = trx.Exec(q, postid)
	if err != nil {
		trx.Rollback()
		log.Printf("failed deleting old author references, err: %s", err)
		return err
	}

	q = `
		INSERT INTO author_post(authorid, postid)
		VALUES (:authorid, :postid)
		ON CONFLICT(authorid, postid)
		DO NOTHING
	`

	r, err := trx.NamedExecContext(ctx, q, author_post)
	if err != nil {
		trx.Rollback()
		return err
	}

	if a, err := r.RowsAffected(); err != nil {
		trx.Rollback()
		return err
	} else if a == 0 {
		trx.Rollback()
		log.Printf("failed updating author_post relationship, err: %s", err)
		return errors.New("failed updating author_post relationship")
	}

	err = trx.Commit()
	if err != nil {
		log.Printf("failed upserting post, could not commit transaction %s", err.Error())
		return errors.New("failed updating post_author relationship")
	}

	return nil
}
