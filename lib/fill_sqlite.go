package lib

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

func FillSqlite(dir string) {
	db, _ := badger.Open(badger.DefaultOptions(dir))
	defer db.Close()
	sdb := OpenSqliteDB()
	defer sdb.Close()
	EnumerateKeysFillSqlitePosts(sdb, db, []byte{17})
	EnumerateKeysFillSqliteUsers(sdb, db, []byte{23})
}

func EnumerateKeysFillSqlitePosts(sdb *sql.DB, db *badger.DB, dbPrefix []byte) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := dbPrefix

		i := 0
		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			val, _ := nodeIterator.Item().ValueCopy(nil)

			post := &PostEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(post)
			InsertPost(sdb, post)
			i++
			if i%1000 == 0 {
				fmt.Println("iteration", i)
			}
		}
		return nil
	})

}
func EnumerateKeysFillSqliteUsers(sdb *sql.DB, db *badger.DB, dbPrefix []byte) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := dbPrefix

		i := 0
		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			val, _ := nodeIterator.Item().ValueCopy(nil)

			profile := &ProfileEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(profile)

			InsertUser(sdb, profile)
			i++
			if i%1000 == 0 {
				fmt.Println("iteration", i)
			}
		}
		return nil
	})

}
