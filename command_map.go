package main

import (
	"fmt"

	internal "github.com/NickLiu-0717/pokedexcli/internal/pokedata"
)

func commandMap(cfg *Config) error {
	areas, err := internal.FetchAreas(cfg.Next)
	if err != nil {
		return fmt.Errorf("error fetching areas: %w", err)
	}
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	cfg.Next = areas.Next
	if areas.Previous == nil {
		cfg.Previous = ""
	} else {
		previous, ok := areas.Previous.(string)
		if !ok {
			fmt.Println("Error: area.Previous is not a string")
		} else {
			cfg.Previous = previous
		}
	}

	return nil

}

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
	} else {
		areas, err := internal.FetchAreas(cfg.Previous)
		if err != nil {
			return fmt.Errorf("error fetching areas: %w", err)
		}
		for _, area := range areas.Results {
			fmt.Println(area.Name)
		}
		cfg.Next = areas.Next
		if areas.Previous == nil {
			cfg.Previous = ""
		} else {
			previous, ok := areas.Previous.(string)
			if !ok {
				fmt.Println("Error: area.Previous is not a string")
			} else {
				cfg.Previous = previous
			}
		}
	}
	return nil
}
