// Task: Write a Go utility that converts os.Args into a slice of structures with fields
// to store the index and value of each command line argument.
// You must determine the structure to use yourself.
package main

import (
	"fmt"
	"os"
)

type newEntry struct {
	Index int
	Value string
}

func argsToSlice(idx int, val string) newEntry {
	return newEntry{Index: idx, Value: val}
}

func main() {

	var newData []newEntry
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	s := arguments[1:]
	for i, item := range s {
		entry := argsToSlice(i, item)
		newData = append(newData, entry)
	}
	fmt.Println(newData)

}
