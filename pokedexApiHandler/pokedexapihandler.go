package pokedexApiHandler

import (
	"encoding/json"
	"net/http"
)

func GetLocationArea(url string) (map[string]any, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return map[string]any{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return map[string]any{}, err
	}
	defer res.Body.Close()

	var location map[string]any

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&location); err != nil {
		return map[string]any{}, err
	}

	return location, nil
}
