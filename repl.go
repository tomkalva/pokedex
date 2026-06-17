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
	pokedex       map[string]pokeapi.PokemonResponse
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		words := cleanInput(input)

		if len(words) < 1 {
			fmt.Print("No input given\n")
			continue
		}
		commandName := words[0]
		args := words[1:]

		cmd, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := cmd.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input:", err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
