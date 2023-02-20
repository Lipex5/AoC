package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number = 0
var sum = 0
var max [3]int

func main() {
	max[0] = 0
	max[1] = 0
	max[2] = 0

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		number, err = strconv.Atoi(fileScanner.Text())
		if err != nil {
			if sum > max[2] {
				max[2] = sum
				rotate_max()
				sum = 0
			} else {
				sum = 0
			}
		} else {
			sum = sum + number
			fmt.Println("Sum: ")
			fmt.Println(sum)
		}
	}

	max_sum := max[0] + max[1] + max[2]

	fmt.Print("Max sum 1: ")
	fmt.Println(max[0])
	fmt.Print("Max sum 2: ")
	fmt.Println(max[1])
	fmt.Print("Max sum 3: ")
	fmt.Println(max[2])
	fmt.Print("Max sum: ")
	fmt.Println(max_sum)

	readFile.Close()
}

func rotate_max() {
	if max[2] > max[1] {
		max[1], max[2] = max[2], max[1]
	}
	if max[1] > max[0] {
		max[0], max[1] = max[1], max[0]
	}
}
