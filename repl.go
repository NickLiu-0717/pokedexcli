package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokecache "github.com/NickLiu-0717/pokedexcli/internal/pokecache"
	pokedata "github.com/NickLiu-0717/pokedexcli/internal/pokedata"
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
			if clninput[0] == "explore" {
				if len(clninput) != 2 {
					fmt.Println("Incorrect Input")
					continue
				} else {
					cfg.Area = clninput[1]
				}
			}
			if clninput[0] == "catch" {
				if len(clninput) != 2 {
					fmt.Println("Incorrect Input")
					continue
				} else {
					cfg.pokemon = clninput[1]
				}
			}
			if clninput[0] == "inspect" {
				if len(clninput) != 2 {
					fmt.Println("Incorrect Input")
					continue
				} else {
					if _, ok := cfg.Pokedex[clninput[1]]; !ok {
						fmt.Println("you have not caught that pokemon")
						continue
					} else {
						cfg.pokemon = clninput[1]
					}
				}
			}
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
		Area:     "",
		pokemon:  "",
		Pokedex:  make(map[string]pokedata.Pokemon),
	}
}

type Config struct {
	Next     string
	Previous string
	cache    *pokecache.Cache
	Area     string
	pokemon  string
	Pokedex  map[string]pokedata.Pokemon
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
		"explore": {
			name:        "explore",
			description: "Displays pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect pokemon information",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}
