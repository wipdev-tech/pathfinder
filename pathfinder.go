package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		printHelp()
		return
	}

	if os.Args[1] == "list" {
		sortPaths := len(os.Args) == 3 && os.Args[2] == "--sort"
		printList(sortPaths)
		return
	}
}

func printHelp() {
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

func printList(sort bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	if sort {
		slices.Sort(paths)
	}

	for _, p := range paths {
		fmt.Println(p)
	}
}
