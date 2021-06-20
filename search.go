package main

import (
	"bcs/lib"
	"fmt"
)

func HandleSearch() {
	query := argMap["query"]
	if query == "" {
		fmt.Println("run with --query=term")
		return
	}
	lib.SearchSqlite(query)
}
