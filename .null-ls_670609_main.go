package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data , err := os.ReadDir("./enron1/")
	if err != nil {
	fmt.Println("Error Reading File", err)
		return
	}

	bag := make(map[string]int)
	tokens := strings.Split(string(data), " ")
	for i := range tokens {
		bag[tokens[i]]++
	}

	for key, value := range bag {
		fmt.Println("Token", key, ":", value)
	}
}
