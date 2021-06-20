package lib

import (
	"database/sql"
	"fmt"

	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/dgraph-io/badger/v3"
)

func InsertPost(db *badger.DB, sdb *sql.DB, post *PostEntry) {
	tx, _ := sdb.Begin()

	body := string(post.Body)
	hash := base58.Encode(post.PostHash.Bytes())
	username := LookupUsername(db, post.PosterPublicKey)

	s := `insert into posts (username, hash, body, created_at) values (?, ?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
	}
	_, e = thing.Exec(username, hash, body, time.Now())
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
