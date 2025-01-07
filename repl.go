package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokecache "github.com/NickLiu-0717/pokedexcli/internal/pokecache"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := NewConfig()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		clninput := cleanInput(input)
		commands := getCommand()
		if cmd, ok := commands[clninput[0]]; ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println("Unknown Error")
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}
	}

}

func cleanInput(text string) []string {
	lowertext := strings.ToLower(text)
	return strings.Fields(lowertext)

}

func NewConfig() *Config {
	return &Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
		cache:    pokecache.NewCache(5 * time.Second),
	}
}

type Config struct {
	Next     string
	Previous string
	cache    *pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config) error
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 areas",
			callback:    commandMapb,
		},
	}
}
