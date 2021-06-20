package lib

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/dgraph-io/badger/v3"
)

func VisualizeSocialGraph(dir string) {
	db, _ := badger.Open(badger.DefaultOptions(dir))
	defer db.Close()
	PrefixFollowerPKIDToFollowedPKID := byte(28)
	PrefixFollowedPKIDToFollowerPKID := byte(29)
	ListFollower2Followed(db, []byte{PrefixFollowerPKIDToFollowedPKID})
	ListFollowed2Follower(db, []byte{PrefixFollowedPKIDToFollowerPKID})

	list := []string{}
	for k, v := range Lookup {
		if v == "404" {
			continue
		}
		list = append(list, k)
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	nodeMap := map[string]int{}

	buff := []string{}
	buff = append(buff, "digraph regexp {")
	for i, item := range list {
		buff = append(buff, fmt.Sprintf(" n%d [label=\"%s\"];", i, Lookup[item]))
		nodeMap[item] = i
	}
	for k, v := range Follower2Followed {
		buff = append(buff, fmt.Sprintf(" n%d -> n%d;", nodeMap[k], nodeMap[v]))
	}
	for k, v := range Followed2Follower {
		buff = append(buff, fmt.Sprintf(" n%d -> n%d;", nodeMap[v], nodeMap[k]))
	}
	buff = append(buff, "}")

	ioutil.WriteFile("clout.gv", []byte(strings.Join(buff, "\n")), 0755)
	fmt.Println("clout.gv file writen.")
}

func ListFollowed2Follower(db *badger.DB, dbPrefix []byte) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := dbPrefix

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			key := nodeIterator.Item().Key()
			followed := key[1:34]
			follower := key[34:]
			HandleFollowed2Follower(db, followed, follower)
		}
		return nil
	})

}
func ListFollower2Followed(db *badger.DB, dbPrefix []byte) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := dbPrefix

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			key := nodeIterator.Item().Key()
			follower := key[1:34]
			followed := key[34:]
			HandleFollower2Followed(db, follower, followed)
		}
		return nil
	})

}
