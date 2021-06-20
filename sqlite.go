package main

import "bcs/lib"

func HandleSqlite() {
	dir := DirCheck()
	if dir == "" {
		return
	}
	lib.CreateSchema()
	lib.FillSqlite(dir)
}
