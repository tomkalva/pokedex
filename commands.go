package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	var commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
	}
	return commands
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for name, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	data, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextURL)
	if err != nil {
		return err
	}
	cfg.nextURL = data.Next
	cfg.previousURL = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	data, err := cfg.pokeapiClient.GetLocationAreas(cfg.previousURL)
	if err != nil {
		return err
	}
	cfg.nextURL = data.Next
	cfg.previousURL = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	return nil
}
