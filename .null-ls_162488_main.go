package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Parsing(path string) map[string]int {
	bag := make(map[string]int)
	var paths []string
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
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

		for _, token := range tokens {
			reg, _ := regexp.Compile("[^a-zA-Z]+")
			result := reg.ReplaceAllString(token, "")
			bag[result]++
		}
	}
	return bag
}

func main() {
	spamBag := make(map[string]int)
	hamBag := make(map[string]int)
	spamBag = Parsing("./enron1/spam/")
	hamBag = Parsing("./enron1/ham/")

		
	totalSpamEmails := 100
	totalEmails := 200
	priorSpam := float64(totalSpamEmails) / float64(totalEmails)



	for key, value := range spamBag {
		fmt.Println("Token", key, ":", value)
	}
}

