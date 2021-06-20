package main

import "bcs/lib"

func HandleSqlite() {
	dir := DirCheck()
	if dir == "" {
		return
	}
	lib.CreateSchema()
	if argMap["testing"] != "" {
		lib.Testing = true
	}
	lib.FillSqlite(dir)
}
