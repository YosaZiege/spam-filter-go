package main

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Dataset struct {
	TotalWords  int
	TotalEmails int
	WordsBag    map[string]int
}

func Parsing(path string) Dataset {
	var dt Dataset
	bag := make(map[string]int)
	dt.WordsBag = bag
	var paths []string
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			paths = append(paths, path)
			dt.TotalEmails++
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
			dt.WordsBag[result]++
			dt.TotalWords++
		}
	}
	return dt
}

func pWordIsSpam(wordCount int, totalSpamWords int, vocabSize int) float64 {
	return float64(wordCount+1) / float64(totalSpamWords+vocabSize)
}

func tokenize(email string) map[string]int {
	tokens := strings.Split(email, " ")
	tokenMap := make(map[string]int)
	for _, token := range tokens {
		reg, _ := regexp.Compile("[^a-zA-Z]+")
		result := reg.ReplaceAllString(token, "")
		tokenMap[result]++
	}
	return tokenMap
}

func classify(email string, spamBag, hamBag Dataset, priorSpam float64, totalSpamWords int, vocabSize int) bool {
	words := tokenize(email)

	spamScore := math.Log(priorSpam)
	hamScore := math.Log(1 - priorSpam)

	for _, word := range words {
		spamScore += math.Log(pWordIsSpam(word, spamBag.TotalWords, vocabSize))
		hamScore += math.Log(pWordIsSpam(word, hamBag.TotalWords, vocabSize))
	}
	fmt.Printf("Spam score: %.4f\n", spamScore)
	fmt.Printf("Ham score:  %.4f\n", hamScore)
	if spamScore > hamScore {
		return true
	} else {
		return false
	}
}

func main() {
	var spamBag Dataset
	var hamBag Dataset
	spamBag = Parsing("./enron1/spam/")
	hamBag = Parsing("./enron1/ham/")

	vocab := make(map[string]bool)
	for w := range spamBag.WordsBag {
		vocab[w] = true
	}
	for w := range hamBag.WordsBag {
		vocab[w] = true
	}
	vocabSize := len(vocab)

	totalSpamEmails := spamBag.TotalEmails
	totalSpamWords := spamBag.TotalWords
	totalEmails := spamBag.TotalEmails + hamBag.TotalEmails

	priorSpam := float64(totalSpamEmails) / float64(totalEmails)

	testEmail := "free viagra offer click here now"
	if classify(testEmail, spamBag, hamBag, priorSpam, totalSpamWords, vocabSize) {
		fmt.Println("SPAM")
	} else {
		fmt.Println("NOT SPAM")
	}
}
