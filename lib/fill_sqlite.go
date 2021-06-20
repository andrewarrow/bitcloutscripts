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
	PrefixPostHashToPostEntry := byte(17)
	EnumerateKeysFillSqlite(sdb, db, []byte{PrefixPostHashToPostEntry})
}

func EnumerateKeysFillSqlite(sdb *sql.DB, db *badger.DB, dbPrefix []byte) {

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
