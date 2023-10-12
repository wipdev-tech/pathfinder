package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/wipdev-tech/pathfinder/internal/pf"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		pf.PrintHelp()
		return
	}

	if os.Args[1] == "list" {
		listArgs := os.Args[2:]
		sortPaths := slices.Contains(listArgs, "--sort")
		searchString := ""

		for i, arg := range listArgs {
			if arg == "--search" {
				searchString = listArgs[i+1]
                break
			}
		}

		pf.PrintList(sortPaths, searchString)
		return
	}

    fmt.Println("Pathfinder: unknown command")
    fmt.Println("Run 'pathfinder help' for info")
}
