package database

import "database/sql"

type Author struct {
	AuthorId  int32
	Username  string
	Firstname string
	Lastname  string
	Email     string
}

type Post struct {
	Body    sql.NullString
	Created sql.NullTime
	PostId  sql.NullInt32
	Title   string
	Updated sql.NullTime
}

type AuthorPost struct {
	ID       sql.NullInt32
	AuthorId int32
	PostId   int32
}

type PostById struct {
	Post
	Firstnames []string
	Lastnames  []string
	Usernames  []string
	Emails     []string
}
