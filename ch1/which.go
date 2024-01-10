package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please provide at least one argument!")
		return
	}

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, file := range arguments[1:] {
		fmt.Printf("Searching for: %s\n", file)
		found := false

		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, file)
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				if mode.IsRegular() {
					if mode&0111 != 0 {
						fmt.Println(fullPath)
						found = true
					}
				}
			}
		}

		if !found {
			fmt.Printf("%s not found\n", file)
		}
	}
}
