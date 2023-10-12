package pf

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

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
