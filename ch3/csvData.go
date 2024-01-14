// Task: Change csvData.go to separate record fields based on
// symbol.

// Second Task: Modify csvData.go to be able to split
// record fields with a symbol, which is specified as a command argument
// lines.
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

func readCSVFilee(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// CSV file read all at once
	// lines data type is [][]string
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

// Change file name for avoid linter errors
func saveCSVFilee(filepath string, delimiter rune) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	// Changing the default field delimiter to tab
	//csvwriter.Comma = '#' // Just change it Task: 1
	csvwriter.Comma = delimiter // Task: 2
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("csvData input output!")
		return
	}

	input := os.Args[1]
	output := os.Args[2]
	delimiter := []rune(os.Args[3])

	if len(delimiter) != 1 {
		fmt.Println("Delimiter must be a single character")
		return
	}

	lines, err := readCSVFilee(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// CSV data is read in columns - each line is a slice
	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
		fmt.Println(temp)
	}

	err = saveCSVFilee(output, delimiter[0])
	if err != nil {
		fmt.Println(err)
		return
	}
}
