package main

import "fmt"

func commandPokedex(cfg *Config) error {
	fmt.Println("Your Pokedex:")
	for name := range cfg.Pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
