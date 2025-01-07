package internal

import (
	"encoding/json"
	"io"
	"net/http"

	pokecache "github.com/NickLiu-0717/pokedexcli/internal/pokecache"
)

func FetchAreas(url string, cache *pokecache.Cache) (Area, error) {
	var area Area
	if data, found := cache.Get(url); found {
		if err := json.Unmarshal(data, &area); err != nil {
			return Area{}, err
		}
		return area, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return Area{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Area{}, err
	}

	cache.Add(url, data)

	defer res.Body.Close()
	if err := json.Unmarshal(data, &area); err != nil {
		return Area{}, err
	}
	return area, nil
}

func FetchLocationPokemon(url string, cache *pokecache.Cache) (Pokeinfo, error) {
	baseurl := "https://pokeapi.co/api/v2/location-area/"
	fullurl := baseurl + url
	var poke Pokeinfo
	if data, found := cache.Get(fullurl); found {
		if err := json.Unmarshal(data, &poke); err != nil {
			return Pokeinfo{}, err
		}
		return poke, nil
	}
	res, err := http.Get(fullurl)
	if err != nil {
		return Pokeinfo{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokeinfo{}, err
	}

	cache.Add(fullurl, data)

	defer res.Body.Close()
	if err := json.Unmarshal(data, &poke); err != nil {
		return Pokeinfo{}, err
	}
	return poke, nil
}

func FetchPokemonInfo(url string, cache *pokecache.Cache) (Pokemon, error) {
	baseurl := "https://pokeapi.co/api/v2/pokemon/"
	fullurl := baseurl + url
	var poke Pokemon
	if data, found := cache.Get(fullurl); found {
		if err := json.Unmarshal(data, &poke); err != nil {
			return Pokemon{}, err
		}
		return poke, nil
	}
	res, err := http.Get(fullurl)
	if err != nil {
		return Pokemon{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	cache.Add(fullurl, data)

	defer res.Body.Close()
	if err := json.Unmarshal(data, &poke); err != nil {
		return Pokemon{}, err
	}
	return poke, nil
}
