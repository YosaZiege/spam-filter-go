package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// ONLY READS ONE DIRECTORY
	// data, err := os.ReadDir("./enron1")
	// if err != nil {
	// 	fmt.Println("Error Reading File", err)
	// 	return
	// }

	bag := make(map[string]int)
	var paths []string
	err := filepath.WalkDir("./enron1/", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			paths = append(paths, path)
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		panic("Error Traversing Directory")
	}
	for _, entry := range paths {
		data, err := os.ReadFile(entry)
		if err != nil {
			panic("This path doesnt work")
		}

		tokens := strings.Split(string(data), " ")
		for i := range tokens {
			bag[tokens[i]]++
		}
	}
	for key, value := range bag {
		fmt.Println("Token", key, ":", value)
	}
}
