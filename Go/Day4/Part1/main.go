package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputfile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)
	validCount := 0
	for scanner.Scan() {
		passphrase := scanner.Text()
		words := strings.Split(passphrase, " ")

		if isValid(words) {
			validCount++
		}
	}
	fmt.Println("There are", validCount, "valid passphrases")
}

func isValid(words []string) bool {
	wm := map[string]string{}

	for _, word := range words {
		if wm[word] == "" {
			wm[word] = "done"
		} else {
			return false
		}
	}

	return true
}
