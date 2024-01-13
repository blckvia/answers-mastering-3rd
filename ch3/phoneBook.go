package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

var CSVFILE = "ch3/phonebook.csv"

var data = []Entry{}
var index map[string]int

func readCSVFile(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		return err
	}

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()

	for _, line := range lines {
		data = append(data, Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		})
	}

	return nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()

	return nil
}

func deleteEntry(key string) error {
	if i, ok := index[key]; !ok {
		return fmt.Errorf("%s not found", key)
	} else {
		data = append(data[:i], data[i+1:]...)
		delete(index, key)
	}

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}

	return nil
}

func createIndex() error {
	index = make(map[string]int)
	for i, v := range data {
		index[v.Tel] = i
	}

	return nil
}

func matchTel(tel string) bool {
	t := []byte(tel)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func initS(N, S, T string) *Entry {
	// Both of them should have a value
	if T == "" || S == "" {
		return nil
	}
	// Give LastAccess a value
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

func insert(pS *Entry) error {
	if _, ok := index[pS.Tel]; ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	data = append(data, *pS)

	_ = createIndex()

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}

	return nil
}

func search(key string) *Entry {
	if i, ok := index[key]; !ok {
		return nil
	} else {
		data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
		return &data[i]
	}
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Creating:", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileInfo, err := os.Stat(CSVFILE)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "is not a regular file")
	}

	err = readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Differentiating between the commands
	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !matchTel(t) {
			fmt.Println("Invalid telephone number")
			return
		}

		temp := initS(arguments[2], arguments[3], t)
		if temp != nil {
			err := insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Invalid telephone number")
			return
		}
		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Invalid telephone number")
			return
		}
		result := search(t)
		if result == nil {
			fmt.Println("Entry not found:", t)
			return
		}
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}
