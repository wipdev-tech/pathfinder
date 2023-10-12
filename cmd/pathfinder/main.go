package main

import (
	"os"

    "github.com/wipdev-tech/pathfinder/internal/pf"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		pf.PrintHelp()
		return
	}

	if os.Args[1] == "list" {
		sortPaths := len(os.Args) == 3 && os.Args[2] == "--sort"
		pf.PrintList(sortPaths)
		return
	}
}
