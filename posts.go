package main

import (
	"bcs/lib"
	"fmt"
)

func DirCheck() string {
	dir := argMap["dir"]
	if dir == "" {
		fmt.Println("run with --dir=/home/name/path/to/badgerdb")
		return ""
	}
	return dir
}

func HandlePosts() {
	dir := DirCheck()
	if dir == "" {
		return
	}
	lib.PrintEveryClout(dir)
}
