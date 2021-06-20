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
create table posts (hash text, body text, created_at datetime);

CREATE UNIQUE INDEX posts_hash_idx
  ON posts (hash);

create table users (bio text, username text, pub58 text, created_at datetime);

CREATE UNIQUE INDEX users_idx
  ON users (pub58);
`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
