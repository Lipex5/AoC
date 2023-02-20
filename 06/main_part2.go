package main

import (
	"bufio"
	"fmt"
	"os"
)

const n_of_chars = 7

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	array := make([]rune, n_of_chars)

	for fileScanner.Scan() {

		// Add first 4 inputs
		for i := 0; i < n_of_chars; i++ {
			array[i] = rune(fileScanner.Text()[i])
		}

		// Loop through all chars of the input
		for i := n_of_chars; i < len(fileScanner.Text()); i++ {
			if check_different(array) {
				fmt.Print("Found at: ")
				fmt.Println(i)
				break
			}
			array = rotate_array(array)
			array[n_of_chars-1] = rune(fileScanner.Text()[i])
		}
	}
}

func rotate_array(arr []rune) []rune {
	for i := 0; i < n_of_chars-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[n_of_chars-1] = rune(' ')

	return arr
}

func check_different(arr []rune) bool {
	for i := 0; i < n_of_chars; i++ {
		for j := 0; j < n_of_chars; j++ {
			if i == j {
			} else if arr[i] == arr[j] {
				return false
			}
		}
	}
	return true
}
