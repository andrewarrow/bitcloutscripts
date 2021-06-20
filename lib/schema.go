package lib

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenSqliteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "clout.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func CreateSchema() {
	db := OpenSqliteDB()
	defer db.Close()

	sqlStmt := `
create table posts (username text, hash text, body text, created_at datetime);

CREATE UNIQUE INDEX posts_hash_idx
  ON posts (hash);

CREATE INDEX posts_username_idx
  ON posts (username);

create table users (bio text, username text, pub58 text, created_at datetime);

CREATE UNIQUE INDEX users_idx
  ON users (pub58);

create table user_follower (followee text, follower text);

CREATE INDEX uf_followee_idx
  ON user_follower (followee);

CREATE INDEX uf_follower_idx
  ON user_follower (follower);
`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
