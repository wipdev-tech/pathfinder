package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		PrintHelp()
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

		PrintList(sortPaths, searchString)
		return
	}

    fmt.Println("Pathfinder: unknown command")
    fmt.Println("Run 'pathfinder help' for info")
}

func PrintHelp() {
	helpLines := []string{
		"Pathfinder - A simple CLI for managing path variables",
		"",
		"Commands:",
		"    help                     Print this message",
		"    list [--sort]            List paths in the PATH variable",
		"                             (--sort orders them alphabettically)",
		"    add <path> [--persist]   Add <path> to the PATH variable",
		"                             (--persist modifies .bashrc to persist the new path)",
	}

	for _, l := range helpLines {
		fmt.Println(l)
	}
}


func PrintList(sort bool, searchStr string) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	if sort {
		slices.Sort(paths)
	}

	for _, p := range paths {
		if strings.Contains(p, searchStr) {
			fmt.Println(p)
		}
	}
}
