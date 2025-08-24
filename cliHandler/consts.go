package cliHandler

const (
	POKEAPI_BASE_URL          = "https://pokeapi.co/api/v2/"
	POKEDEX_LOCATION_AREA_URL = "https://pokeapi.co/api/v2/location-area/"
	POKEDEX_POKEMON_URL       = "https://pokeapi.co/api/v2/pokemon/"
)

const (
	POKEDEX_LOCATION_AREA_LIMIT           = 20
	POKEDEX_LOCATION_AREA_MAX_LOCATION_ID = 64
	POKEDEX_POKEMON_CATCH_DIFFICULTY      = 100
)

const (
	REPL_EXIT = iota
	REPL_HELP
	REPL_MAP
	REPL_MAPB
)
