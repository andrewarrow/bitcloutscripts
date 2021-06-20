package lib

import (
	"bytes"
	"database/sql"
	"encoding/gob"

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

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			val, _ := nodeIterator.Item().ValueCopy(nil)

			post := &PostEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(post)
			InsertPost(sdb, post)
		}
		return nil
	})

}
