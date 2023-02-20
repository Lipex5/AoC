package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		temp := strings.Split(line, ",")
		tempA := strings.Split(temp[0], "-")
		tempB := strings.Split(temp[1], "-")
		A := make([]int, 2)
		B := make([]int, 2)
		A[0], _ = strconv.Atoi(tempA[0])
		A[1], _ = strconv.Atoi(tempA[1])
		B[0], _ = strconv.Atoi(tempB[0])
		B[1], _ = strconv.Atoi(tempB[1])

		if A[0] < B[0] && A[1] < B[0] {
			continue
		}
		if B[0] < A[0] && B[1] < A[0] {
			continue
		}
		sum++
	}
	fmt.Println(sum)
}
