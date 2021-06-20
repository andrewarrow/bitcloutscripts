package main

import (
	"bcs/files"
	"bcs/lib"
	"fmt"
	"strings"
)

func DirCheck() string {
	dir := argMap["dir"]
	if dir == "" {
		fmt.Println("run with --dir=/home/name/path/to/badgerdb")
		return ""
	}
	if strings.HasPrefix(dir, "~") {
		home := files.UserHomeDir()
		return home + dir[1:]
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
