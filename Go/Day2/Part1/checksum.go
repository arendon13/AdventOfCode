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
	
	digitsToSum := getDigitsToSum(grid)
	
	sum := checksum(digitsToSum)

  fmt.Println("Value for checksum is", sum)
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

func getDigitsToSum(grid [][]int) []int {
	digitsToSum := []int{}
	
	for _, row := range grid {
		var min, max int
		for _, rownum := range row {
			if min == 0 {
				min = rownum
			}
			if rownum > max {
				max = rownum
			}
			if rownum < min {
				min = rownum
			}
		}
		digitsToSum = append(digitsToSum, max-min)
	}

  return digitsToSum
}

func checksum(digitsToSum []int) int {
	var sum int

	for i := range digitsToSum {
		sum += digitsToSum[i]
	}

	return sum
}
