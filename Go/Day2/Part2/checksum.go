package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := getGrid()

	rowDivNums := getDivisionNumbers(grid)

	sum := checksum(rowDivNums)

	fmt.Println("Value for checksum is", sum)
}

func getGrid() [][]int {
	inputfile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)
	grid := [][]int{}

	for scanner.Scan() {
		input := scanner.Text()
		inputNums := strings.Split(input, "\t")
		numbers := getNumbers(inputNums)
		grid = append(grid, numbers)
	}

	return grid
}

func getNumbers(stringNums []string) []int {
	nums := make([]int, len(stringNums))
	for i := range stringNums {
		parsedNum, err := strconv.Atoi(stringNums[i])
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = parsedNum
	}

	return nums
}

func getDivisionNumbers(grid [][]int) []int {
	rowDivNums := []int{}

	for _, row := range grid {
		div := getDivision(row)
		rowDivNums = append(rowDivNums, div)
	}

	return rowDivNums
}

func getDivision(row []int) int {
	for i := range row {
		for j := range row {
			mod := row[i] % row[j]
			if i != j && mod == 0 {
				return row[i] / row[j]
			}
		}
	}

	return 0
}

func checksum(digitsToSum []int) int {
	var sum int

	for i := range digitsToSum {
		sum += digitsToSum[i]
	}

	return sum
}
