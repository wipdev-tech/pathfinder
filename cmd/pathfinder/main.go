package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		PrintHelp()
		return
	}

	switch os.Args[1] {
	case "help":
		PrintHelp()

	case "list":
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

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Pathfinder: no path provided")
			return
		}

        checkDotProfile()
		// addPath(os.Args[2])

	default:
		fmt.Println("Pathfinder: unknown command")
		fmt.Println("Run 'pathfinder help' for info")
	}
}

func PrintHelp() {
	helpLines := []string{
		"Pathfinder - A simple CLI for managing path variables",
		"",
		"Commands:",
		"    help                     Print this message",
		"    list [--sort]            List paths in the PATH variable",
		"         [--search string]",
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

func checkDotProfile() {
	fr, err := os.OpenFile(
		os.Getenv("HOME")+"/.profile",
		os.O_RDONLY, 0644,
	)
	if err != nil {
		fmt.Println(err)
	}
	defer fr.Close()

	scanner := bufio.NewScanner(fr)
    fmt.Println("before:")
	for scanner.Scan() {
		line := scanner.Text()
        fmt.Println(line)
	}

}

// func addPath(newPath string) {
// 	fw, err := os.OpenFile(
// 		os.Getenv("HOME")+"/Downloads/hamada.txt",
// 		os.O_WRONLY|os.O_APPEND, 0644,
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer fw.Close()
//
//     _, err = fw.WriteString("a7lamesa3lek\n")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
