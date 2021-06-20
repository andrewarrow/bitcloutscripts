package lib

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/dgraph-io/badger/v3"
)

var Testing bool

func FillSqlite(dir string) {
	db, _ := badger.Open(badger.DefaultOptions(dir))
	defer db.Close()
	sdb := OpenSqliteDB()
	defer sdb.Close()
	EnumerateKeysFillSqlitePosts(sdb, db, []byte{17})
	EnumerateKeysFillSqliteUsers(sdb, db, []byte{23})
	EnumerateKeysFillSqliteFollows(sdb, db, []byte{29})

}
func EnumerateKeysFillSqliteFollows(sdb *sql.DB, db *badger.DB, dbPrefix []byte) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := dbPrefix

		i := 0
		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			key := nodeIterator.Item().Key()
			follower := key[1:34]
			followed := key[34:]
			InsertFollowee(sdb, base58.Encode(followed), base58.Encode(follower))
			i++
			if i%1000 == 0 {
				fmt.Println("iteration", i)
				if Testing {
					break
				}
			}
		}
		return nil
	})

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
			InsertPost(db, sdb, post)
			i++
			if i%1000 == 0 {
				fmt.Println("iteration", i)
				if Testing {
					break
				}
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
				if Testing {
					break
				}
			}
		}
		return nil
	})

}
