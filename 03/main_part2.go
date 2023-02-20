package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var letter rune
	sum := 0
	var text [3]string
	i := 0

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text[i] = fileScanner.Text()

		if i == 2 {
			for _, char := range text[0] {
				for _, comp := range text[1] {
					for _, comp2 := range text[2] {
						if char == comp && char == comp2 {
							letter = char
							i = -1
						}
					}
				}
			}
			sum += get_score(letter)
		}
		i++

		fmt.Println(fileScanner.Text())
		fmt.Println(text[0])
		fmt.Println(text[1])
		fmt.Println(string(letter))
		fmt.Println(int(letter - ('A' - 1)))
		fmt.Println(sum)
	}
}

func get_score(letter rune) int {
	// Lower case
	if letter > ('a' - 1) {
		fmt.Print("Score: ")
		fmt.Println(int(letter - ('a' - 1)))
		return (int(letter - ('a' - 1)))
	} else {
		fmt.Print("Score: ")
		fmt.Println(int(letter-('A'-1)) + 26)
		return (int(letter-('A'-1)) + 26)
	}
}
