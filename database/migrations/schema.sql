CREATE TABLE author (
    authorid SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    email VARCHAR(50) NOT NULL,
    UNIQUE(email)
);

CREATE TABLE post (
    body VARCHAR(5000),
    created TIMESTAMP DEFAULT current_timestamp,
    postid SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    updated TIMESTAMP 
);

CREATE TABLE author_post (
    id SERIAL PRIMARY KEY,
    authorid INT NOT NULL,
    postid INT NOT NULL,
    UNIQUE(authorid, postid),
    CONSTRAINT fk_postid FOREIGN KEY(postid) REFERENCES post(postid),
    CONSTRAINT fk_authorid FOREIGN KEY(authorid) REFERENCES author(authorid) ON DELETE CASCADE
);