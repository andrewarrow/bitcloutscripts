package lib

import (
	"fmt"
	"strconv"
)

func SearchSqliteFollow(tab, s, degrees string) {
	pub58 := SearchSqliteUsername(s)
	db := OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select follower from user_follower where followee = '" + pub58 + "'")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	limit, _ := strconv.Atoi(degrees)
	tabSize := len(tab) / 2

	for rows.Next() {
		var follower string
		rows.Scan(&follower)
		username := SearchSqlitePub58(follower)
		fmt.Printf("%s%s\n", tab, username)

		if tabSize+1 < limit {
			SearchSqliteFollow(tab+"  ", username, degrees)
		}
	}
}
func SearchSqlitePosts(s string) {
	db := OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select username,body from posts where body like '%" + s + "%'")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var body string
		var username string
		rows.Scan(&username, &body)
		fmt.Println(username, body)
	}
}
func SearchSqliteUsers(s string) {
	db := OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select username,bio from users where (username like '%" + s + "%') or (bio like '%" + s + "%')")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		var bio string
		rows.Scan(&username, &bio)
		fmt.Println(username, bio)
	}
}
func SearchSqliteUsername(s string) string {
	db := OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select pub58 from users where username='" + s + "'")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		var pub58 string
		rows.Scan(&pub58)
		return pub58
	}

	return ""
}
func SearchSqlitePub58(s string) string {
	db := OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select username from users where pub58='" + s + "'")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		rows.Scan(&username)
		return username
	}

	return ""
}
