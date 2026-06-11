package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		words := cleanInput(input)
		if len(words) < 1 {
			fmt.Print("No input given\n")
			continue
		}
		firstWord := words[0]
		fmt.Printf("Your command was: %v\n", firstWord)
	}
}
