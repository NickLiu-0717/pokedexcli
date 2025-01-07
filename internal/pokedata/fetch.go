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
