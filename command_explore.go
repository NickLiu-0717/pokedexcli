package main

import (
	"fmt"

	internal "github.com/NickLiu-0717/pokedexcli/internal/pokedata"
)

func commandExplore(cfg *Config) error {
	pokes, err := internal.FetchLocationPokemon(cfg.Area, cfg.cache)
	if err != nil {
		return fmt.Errorf("error fetching pokes: %w", err)
	}
	fmt.Println("Found Pokemon:")
	for _, poke := range pokes.PokemonEncounters {
		fmt.Printf(" - %s\n", poke.Pokemon.Name)
	}
	return nil
}
