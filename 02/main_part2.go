// A - Rock  		1
// B - Paper  		2
// C - Scissors	  	3
// Win Z -		6
// Draw Y - 	3
// Lose X - 	0

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	points := 0
	win := 6
	draw := 3

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		output := strings.Split(fileScanner.Text(), " ")
		switch output[1] {
		case "X":
			switch output[0] {
			case "A":
				points += 3

			case "B":
				points += 1

			case "C":
				points += 2
			}
		case "Y":
			points += draw
			switch output[0] {
			case "A":
				points += 1

			case "B":
				points += 2

			case "C":
				points += 3
			}
		case "Z":
			points += win
			switch output[0] {
			case "A":
				points += 2

			case "B":
				points += 3

			case "C":
				points += 1
			}
		}
	}
	fmt.Println(points)
}
