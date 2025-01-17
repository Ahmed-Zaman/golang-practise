package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	ls := strings.Join(listFiles("testdata"), " ")
	fmt.Println(ls)
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := os.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
