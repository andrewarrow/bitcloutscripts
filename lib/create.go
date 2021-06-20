package lib

import (
	"database/sql"
	"fmt"

	"time"

	"github.com/btcsuite/btcutil/base58"
)

func InsertPost(db *sql.DB, post *PostEntry) {
	tx, _ := db.Begin()

	body := string(post.Body)
	hash := base58.Encode(post.PostHash.Bytes())

	s := `insert into posts (hash, body, created_at) values (?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(hash, body, time.Now())
	if e != nil {
		fmt.Println(e)
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
func InsertUser(db *sql.DB, profile *ProfileEntry) {
	tx, _ := db.Begin()

	pub58 := base58.Encode(profile.PublicKey)
	username := string(profile.Username)
	bio := string(profile.Description)

	s := `insert into users (bio, username, pub58, created_at) values (?, ?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(bio, username, pub58, time.Now())
	if e != nil {
		fmt.Println(e)
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}
