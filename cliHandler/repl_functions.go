package cliHandler

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dmitriy-zverev/pokedex-cli/pokecache"
	"github.com/dmitriy-zverev/pokedex-cli/pokedexApiHandler"
)

func commandExit(config *Config, cache *pokecache.Cache) error {
	updateConfig(REPL_EXIT, config)

	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, cache *pokecache.Cache) error {
	updateConfig(REPL_HELP, config)

	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for key, value := range SupportedFunctions {
		fmt.Printf("%s: %s\n", key, value.Description)
	}

	return nil
}

func commandMap(config *Config, cache *pokecache.Cache) error {
	updateConfig(REPL_MAP, config)

	splittedNextUrl := strings.Split(config.Next, "/")
	nextId, _ := strconv.Atoi(splittedNextUrl[len(splittedNextUrl)-1])

	for i := range POKEDEX_LOCATION_AREA_LIMIT {
		currentId := nextId + i
		if currentId > POKEDEX_LOCATION_AREA_MAX_LOCATION_ID {
			return errors.New("you've scrolled through all of the available locatons")
		}

		fullUrl := POKEDEX_LOCATION_AREA_URL + fmt.Sprint(currentId)
		locationArea, err := pokedexApiHandler.GetLocationArea(fullUrl, cache)
		if err != nil {
			return err
		}

		fmt.Println(locationArea["name"])
	}

	config.Previous = POKEDEX_LOCATION_AREA_URL + fmt.Sprint(nextId+POKEDEX_LOCATION_AREA_LIMIT-1)
	config.Next = POKEDEX_LOCATION_AREA_URL + fmt.Sprint(nextId+POKEDEX_LOCATION_AREA_LIMIT)

	return nil
}

func commandMapb(config *Config, cache *pokecache.Cache) error {
	updateConfig(REPL_MAP, config)

	splittedNextUrl := strings.Split(config.Next, "/")
	nextId, _ := strconv.Atoi(splittedNextUrl[len(splittedNextUrl)-1])

	if nextId <= POKEDEX_LOCATION_AREA_LIMIT*2 {
		return errors.New("you are already on the first page")
	}

	for i := range POKEDEX_LOCATION_AREA_LIMIT {
		currentId := nextId + i
		if currentId > POKEDEX_LOCATION_AREA_MAX_LOCATION_ID {
			return errors.New("you've scrolled through all of the available locatons")
		}

		fullUrl := POKEDEX_LOCATION_AREA_URL + fmt.Sprint(currentId-POKEDEX_LOCATION_AREA_LIMIT*2)
		locationArea, err := pokedexApiHandler.GetLocationArea(fullUrl, cache)
		if err != nil {
			return err
		}

		fmt.Println(locationArea["name"])
	}

	config.Previous = POKEDEX_LOCATION_AREA_URL + fmt.Sprint(nextId-POKEDEX_LOCATION_AREA_LIMIT*2)
	config.Next = POKEDEX_LOCATION_AREA_URL + fmt.Sprint(nextId-POKEDEX_LOCATION_AREA_LIMIT*2+1)

	return nil
}
