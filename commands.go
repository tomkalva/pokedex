package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
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
		"explore": {
			name:        "explore",
			description: "Lists all Pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon and add them to the Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Prints the name, height, weight, stats and type(s) of a Pokemon",
			callback:    commandInspect,
		},
	}
	return commands
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for name, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, args []string) error {
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

func commandMapb(cfg *config, args []string) error {
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

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("location area name required")
	}

	areaName := args[0]

	fmt.Printf("Exploring %s...\n", areaName)

	area, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Pokemon name required")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemonStats(pokemonName)
	if err != nil {
		return err
	}
	_, alreadyCaught := cfg.pokedex[pokemonName]
	if alreadyCaught {
		fmt.Printf("%s is already caught!\n", pokemonName)
	} else {
		fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
		roll := rand.Intn(pokemon.BaseExperience)
		if roll < 100 {
			fmt.Printf("%s was caught!\n", pokemonName)
			cfg.pokedex[pokemonName] = pokemon
		} else {
			fmt.Printf("%s escaped!\n", pokemonName)
		}
	}
	return nil
}

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Pokemon name required")
	}

	pokemonName := args[0]

	_, isCaught := cfg.pokedex[pokemonName]
	if isCaught == false {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %v\n", cfg.pokedex[pokemonName].Name)
		fmt.Printf("Height: %v\n", cfg.pokedex[pokemonName].Height)
		fmt.Printf("Weight: %v\n", cfg.pokedex[pokemonName].Weight)

		fmt.Println("Stats:")
		for _, s := range cfg.pokedex[pokemonName].Stats {
			fmt.Printf("  -%v: %v\n", s.Stat.Name, s.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range cfg.pokedex[pokemonName].Types {
			fmt.Printf("  - %v\n", t.Type.Name)
		}
	}

	return nil
}
