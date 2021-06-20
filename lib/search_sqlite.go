package lib

import (
	"fmt"
)

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
