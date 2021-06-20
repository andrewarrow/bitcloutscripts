package main

import (
	"bcs/lib"
	"fmt"
)

func HandlePosts() {
	dir := argMap["dir"]
	if dir == "" {
		fmt.Println("run with --dir=/home/name/path/to/badgerdb")
		return
	}
	lib.PrintEveryClout(dir)
}
