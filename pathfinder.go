package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 || os.Args[1] == "help" {
        printHelp()
    }
}

func printHelp() {
    helpLines := []string{
        "Pathfinder - A simple CLI for managing path variables",
        "",
        "Commands:",
        "    help                     Print this message",
        "    list                     List paths in the PATH variable",
        "    add <path> [--persist]   Add <path> to the PATH variable",
        "                             (--persist modifies .bashrc to persist the new path)",
    }

    for _, l := range helpLines {
        fmt.Println(l)
    }
}
