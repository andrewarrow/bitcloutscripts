package lib

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

func PrintEveryClout(dir string) {
	db, _ := badger.Open(badger.DefaultOptions(dir))
	defer db.Close()
	PrefixPostHashToPostEntry := byte(17)
	EnumerateKeysForPrefix(db, []byte{PrefixPostHashToPostEntry})
}

func EnumerateKeysForPrefix(db *badger.DB, dbPrefix []byte) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := dbPrefix

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			val, _ := nodeIterator.Item().ValueCopy(nil)

			post := &PostEntry{}
			gob.NewDecoder(bytes.NewReader(val)).Decode(post)
			fmt.Println(string(post.Body))
		}
		return nil
	})

}
