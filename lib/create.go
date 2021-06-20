package lib

import (
	"database/sql"
	"fmt"

	"time"
)

func InsertPost(db *sql.DB, post *PostEntry) {
	tx, _ := db.Begin()

	body := string(post.Body)

	s := `insert into posts (body, created_at) values (?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(body, time.Now())
	if e != nil {
		fmt.Println(e)
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
