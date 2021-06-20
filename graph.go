package main

import "bcs/lib"

func HandleGraph() {
	dir := DirCheck()
	if dir == "" {
		return
	}
	lib.VisualizeSocialGraph(dir)
}
