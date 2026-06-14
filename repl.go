package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
}

func startRepl(cfg *config) {
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

		cmd, ok := getCommands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
