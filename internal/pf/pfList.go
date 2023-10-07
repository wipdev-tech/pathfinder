package pf
import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func PfList(sort bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	if sort {
		slices.Sort(paths)
	}

	for _, p := range paths {
		fmt.Println(p)
	}
}
