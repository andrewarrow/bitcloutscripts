package main

import (
	"bcs/args"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  bcs posts       # print all posts")
	fmt.Println("  bcs graph       # make clout.gv graph file")
	fmt.Println("")
}

var argMap map[string]string

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]
	argMap = args.ToMap()

	if command == "posts" {
		HandlePosts()
	} else if command == "graph" {
		HandleGraph()
	} else {
		fmt.Println("not a command yet")
	}

}
