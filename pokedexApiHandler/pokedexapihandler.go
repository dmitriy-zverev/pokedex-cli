package pokedexApiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dmitriy-zverev/pokedex-cli/pokecache"
)

func GetPokemonData(url string, cache *pokecache.Cache) (map[string]any, error) {
	if cacheData, ok := cache.Get(url); ok {
		var data map[string]any

		if err := json.Unmarshal(cacheData, &data); err != nil {
			return map[string]any{}, err
		}

		return data, nil
	}

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

	var data map[string]any

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&data); err != nil {
		return map[string]any{}, err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return map[string]any{}, err
	}
	cache.Add(url, jsonData)

	return data, nil
}
