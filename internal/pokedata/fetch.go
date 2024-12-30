package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchAreas(url string) (Area, error) {
	res, err := http.Get(url)
	if err != nil {
		return Area{}, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Area{}, err
	}
	defer res.Body.Close()
	var area Area
	if err := json.Unmarshal(data, &area); err != nil {
		return Area{}, err
	}
	return area, nil
}
