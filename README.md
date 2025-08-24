# Your favourite Pokedex CLI
A Pokedex is just a make-believe device that lets us look up information about Pokemon - things like their name, type, and stats

## How to use
Here's the list of commands and what they do:

`help` - displays a help message with all of the commands
```
Pokedex > help
```
```
Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Displays 20 next locations
mapb: Displays 20 previous locations
explore: Expores the location area and returns pokemons at this area
catch: Catches a pokemon—or not
inspect: Inspects a pokemon from your pokedex
pokedex: Lists all of the caught pokemons
exit: Exit the Pokedex
```

`exit` - exits the tool
```
Pokedex > exit
```
```
Closing the Pokedex... Goodbye!
```

`map` - lists the next location area with autopagination on repeat command
```
Pokedex > map
```
```
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

`mapb` - lists the previous location area like `map`
```
Pokedex > mapb
```
```
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

`explore` - explores the given location area and lists all of the pokemons there
```
Pokedex > explore canalave-city-area
```
```
Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - staryu
 - magikarp
 - gyarados
 - wingull
 - pelipper
 - shellos
 - gastrodon
 - finneon
 - lumineon
```

`catch` - catches given pokemon based on randomness
```
Pokedex > catch pikachu
```
```
Throwing a Pokeball at pikachu...
pikachu was caught!
You may now inspect it with the inspect pikachu!
```

`inspect` - lists the information about your caught pokemon
```
Pokedex > inspect pikachu
```
```
Name: pikachu
Height: 4
Weight: 60
Stats:
        hp: 35
        attack: 55
        defense: 40
        special-attack: 50
        special-defense: 50
        speed: 90
Types:
         electric
```

`pokedex` - lists all of the caught pokemons
```
Pokedex > pokedex
```
```
Your Pokedex:
        - pikachu
```