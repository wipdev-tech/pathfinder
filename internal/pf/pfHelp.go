package pf

import "fmt"

func PfHelp() {
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
