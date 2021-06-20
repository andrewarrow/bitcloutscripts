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
	table := "posts"

	if argMap["table"] != "" {
		table = argMap["table"]
	}

	degrees := argMap["degrees"]
	if degrees == "" {
		degrees = "2"
	}

	if table == "posts" {
		lib.SearchSqlitePosts(query)
	} else if table == "users" {
		lib.SearchSqliteUsers(query)
	} else if table == "follow" {
		lib.SearchSqliteFollow("", query, degrees)
	}
}
