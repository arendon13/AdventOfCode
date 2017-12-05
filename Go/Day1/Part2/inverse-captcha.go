package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"log"
	"strings"
)

func main() {
	// Read file
	inputfile, err := os.Open("./input.txt")
	if err != nil {
    log.Fatal(err)
	}
	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)
  for scanner.Scan() {
		// Get line text
		input := scanner.Text()
		// Receiver an array of strings
		inputNums := strings.Split(input, "")
		// Convert into an array of integers
		numbers := getNumbers(inputNums)
		// Calculate the sum
		inverseSum := calcInverseSum(numbers)

		fmt.Println("The sum of captcha number", input, "is", inverseSum)
	}

  if err := scanner.Err(); err != nil {
		log.Fatal(err)
  }

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

func calcInverseSum(nums []int) int {
	var sum int
	numsSize := len(nums)
	steps := numsSize / 2
	for i := range nums {
		indexCompared := getIndex(numsSize - 1, steps, i)
		if nums[indexCompared] == nums[i] {
			sum += nums[i]
		}
	}

	return sum
}

// index required for the circular list
func getIndex(maxIndex int, steps int, curIndex int) int {
	var realIndex int
	misleadingIndex := steps + curIndex
	if misleadingIndex > maxIndex {
		realIndex = misleadingIndex - (maxIndex + 1)
	} else {
		realIndex = misleadingIndex
	}

	return realIndex
}