package lib

import (
	"fmt"
)

func SearchSqlite(s string) {
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
