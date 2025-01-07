package main

import (
	"fmt"
	"math/rand"

	internal "github.com/NickLiu-0717/pokedexcli/internal/pokedata"
)

func commandCatch(cfg *Config) error {
	pokemon, err := internal.FetchPokemonInfo(cfg.pokemon, cfg.cache)
	if err != nil {
		return fmt.Errorf("error fetching pokemons: %w", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", cfg.pokemon)
	base := float64(pokemon.BaseExperience) / 100.0
	randval := rand.Float64() * 4
	if base < randval {
		fmt.Printf("%s was cautht!\n", cfg.pokemon)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.Pokedex[cfg.pokemon] = pokemon
	} else {
		fmt.Printf("%s escaped\n", cfg.pokemon)
	}
	return nil
}
