package cliHandler

import (
	"errors"
	"fmt"
	"math/rand/v2"
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
		locationArea, err := pokedexApiHandler.GetPokemonData(fullUrl, cache)
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
		locationArea, err := pokedexApiHandler.GetPokemonData(fullUrl, cache)
		if err != nil {
			return err
		}

		fmt.Println(locationArea["name"])
	}

	config.Previous = POKEDEX_LOCATION_AREA_URL + fmt.Sprint(nextId-POKEDEX_LOCATION_AREA_LIMIT*2)
	config.Next = POKEDEX_LOCATION_AREA_URL + fmt.Sprint(nextId-POKEDEX_LOCATION_AREA_LIMIT*2+1)

	return nil
}

func commandExplore(config *Config, cache *pokecache.Cache) error {
	if len(config.FullCommand) < 2 {
		return errors.New("cannot explore an empty area")
	}

	locationArea := config.FullCommand[1]
	locationAreaUrl := POKEDEX_LOCATION_AREA_URL + locationArea

	fmt.Printf("Exploring %s...\n", locationArea)

	location, err := pokedexApiHandler.GetPokemonData(locationAreaUrl, cache)
	if err != nil {
		fmt.Println("Location not found")
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, value := range location["pokemon_encounters"].([]any) {
		for i, data := range value.(map[string]any) {
			if i == "pokemon" {
				fmt.Println(" -", data.(map[string]any)["name"])
			}
		}
	}

	return nil
}

func commandCatch(config *Config, cache *pokecache.Cache) error {
	if len(config.FullCommand) < 2 {
		return errors.New("cannot catch nothing")
	}

	pokemon := config.FullCommand[1]
	pokemonUrl := POKEDEX_POKEMON_URL + pokemon

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	pokemonData, err := pokedexApiHandler.GetPokemonData(pokemonUrl, cache)
	if err != nil {
		fmt.Println("Pokemon do not exist")
		return err
	}

	pokemonBaseExperience := int(pokemonData["base_experience"].(float64))

	if pokemonBaseExperience-rand.IntN(pokemonBaseExperience) < POKEDEX_POKEMON_CATCH_DIFFICULTY {
		types := make([]string, 0)

		for _, item := range pokemonData["types"].([]any) {
			types = append(types, item.(map[string]any)["type"].(map[string]any)["name"].(string))
		}

		pokemonStruct := Pokemon{
			Name:           pokemonData["name"].(string),
			Height:         int(pokemonData["height"].(float64)),
			Weight:         int(pokemonData["weight"].(float64)),
			HP:             int(pokemonData["stats"].([]any)[0].(map[string]any)["base_stat"].(float64)),
			Attack:         int(pokemonData["stats"].([]any)[1].(map[string]any)["base_stat"].(float64)),
			Defense:        int(pokemonData["stats"].([]any)[2].(map[string]any)["base_stat"].(float64)),
			SpecialAttack:  int(pokemonData["stats"].([]any)[3].(map[string]any)["base_stat"].(float64)),
			SpecialDefense: int(pokemonData["stats"].([]any)[4].(map[string]any)["base_stat"].(float64)),
			Speed:          int(pokemonData["stats"].([]any)[5].(map[string]any)["base_stat"].(float64)),
			Types:          types,
		}
		config.Pokedex[pokemonStruct.Name] = pokemonStruct

		fmt.Printf("%s was caught!\n", pokemon)
		fmt.Printf("You may now inspect it with the inspect %s!\n", pokemon)
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemon)
	return nil
}

func commandInspect(config *Config, cache *pokecache.Cache) error {
	if len(config.FullCommand) < 2 {
		return errors.New("cannot inspect nothing")
	}

	pokemon := config.FullCommand[1]

	if pokemonData, ok := config.Pokedex[pokemon]; ok {
		fmt.Println("Name:", pokemonData.Name)
		fmt.Println("Height:", pokemonData.Height)
		fmt.Println("Weight:", pokemonData.Weight)
		fmt.Println("Stats:")
		fmt.Println("\thp:", pokemonData.HP)
		fmt.Println("\tattack:", pokemonData.Attack)
		fmt.Println("\tdefense:", pokemonData.Defense)
		fmt.Println("\tspecial-attack:", pokemonData.SpecialAttack)
		fmt.Println("\tspecial-defense:", pokemonData.SpecialDefense)
		fmt.Println("\tspeed:", pokemonData.Speed)
		fmt.Println("Types:")

		for _, t := range pokemonData.Types {
			fmt.Println("\t", t)
		}

		return nil
	}

	fmt.Println("You have not caught this pokemon yet")

	return nil
}

func commandPokedex(config *Config, cache *pokecache.Cache) error {
	fmt.Println("Your Pokedex:")

	if len(config.Pokedex) < 1 {
		fmt.Println("You have not caught any pokemons")
		return nil
	}

	for pokemon := range config.Pokedex {
		fmt.Println("\t-", pokemon)
	}

	return nil
}
