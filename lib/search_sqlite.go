package lib

import (
	"fmt"
)

func SearchSqliteFollow(s string) {
	pub58 := SearchSqliteUsername(s)
	fmt.Println(pub58)
	db := OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select follower from user_follower where followee = '" + pub58 + "'")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var follower string
		rows.Scan(&follower)
		fmt.Println(follower)
	}
}
func SearchSqlitePosts(s string) {
	db := OpenSqliteDB()
	defer db.Close()
	rows, err := db.Query("select body from posts where body like '%" + s + "%'")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var body string
		rows.Scan(&body)
		fmt.Println(body)
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
