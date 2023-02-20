package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	array := make(map[int][]string)
	initcount := 7
	lines := 10

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		// Get arrays from the first part of the input
		if initcount >= 0 {
			j := 0
			for i := 1; i < 35; i += 4 {
				if rune(fileScanner.Text()[i]) != rune(32) {
					array[j] = append(array[j], string(fileScanner.Text()[i]))
				}
				j++
			}
			initcount--
		}

		// Reverse the arrays
		if lines == 1 {
			for i := range array {
				reverse(array[i])
			}
		}

		// Data handling
		if lines == 0 {
			split := strings.Split(fileScanner.Text(), " ")
			qty, _ := strconv.Atoi(split[1])
			from, _ := strconv.Atoi(split[3])
			to, _ := strconv.Atoi(split[5])
			from -= 1
			to -= 1
			fmt.Println(qty)
			fmt.Println(from)
			fmt.Println(to)

			for i := 0; i < qty; i++ {
				if len(array[from]) > 0 {
					array[to] = append(array[to], array[from][len(array[from])-qty])
					array[from] = pop(array[from])
				}
			}
		}

		if lines > 0 {
			lines--
		}
		fmt.Println(array)
	}
	fmt.Println(array)
}

// Pops the last item of the slice
// Must be inputted a slice of ints
func pop(arr []string) []string {
	if len(arr) > 0 {
		return arr[0 : len(arr)-1]
	} else {
		return arr
	}
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
