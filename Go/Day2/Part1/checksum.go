package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	"strings"
)

func main() {
	inputfile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		input := scanner.Text()
		inputNums := strings.Split(input, "\t")
		fmt.Println(inputNums)
	}
}


// func getNumbers(stringNums []string) []int {
// 	nums := make([]int, len(stringNums))
// 	for i := range nums {
// 		parsedNum, err =: str
// 	}
// }