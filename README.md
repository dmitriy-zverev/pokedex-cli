# PokéDex CLI Explorer

A powerful command-line Pokédex application that brings the world of Pokémon to your terminal. Explore locations, catch Pokémon, and build your personal collection with this interactive CLI tool powered by the PokéAPI.

## Description

PokéDex CLI Explorer is an immersive terminal-based application that simulates the classic Pokédex experience. Navigate through different Pokémon locations, discover wild Pokémon in their natural habitats, attempt to catch them, and manage your growing collection—all from the comfort of your command line.

Built with Go, this application features intelligent caching for optimal performance, a clean REPL (Read-Eval-Print Loop) interface, and comprehensive Pokémon data integration through the PokéAPI.

## Why?

**The Problem:** Traditional Pokédex applications are often GUI-heavy and lack the simplicity and efficiency that developers and CLI enthusiasts appreciate. There's a gap in the market for a lightweight, terminal-based Pokémon exploration tool.

**The Solution:** PokéDex CLI Explorer bridges this gap by providing:
- **Developer-Friendly Interface**: Clean command-line experience that feels natural for terminal users
- **Performance Optimization**: Built-in caching system reduces API calls and improves response times
- **Educational Value**: Demonstrates Go programming best practices, API integration, and CLI application design
- **Nostalgia Factor**: Recreates the classic Pokédex experience in a modern, efficient format

**Goals:**
- Provide an engaging way to explore Pokémon data through the terminal
- Showcase clean Go architecture and modular design patterns
- Demonstrate real-world API integration with caching strategies
- Create an educational tool for learning CLI application development

## Quick Start

### Prerequisites
- Go 1.24.5 or higher
- Internet connection (for PokéAPI access)

### Installation & Running

1. **Clone the repository:**
   ```bash
   git clone https://github.com/dmitriy-zverev/pokedex-cli.git
   cd pokedex-cli
   ```

2. __Build and run:__

   ```bash
   go build -o pokedex-cli
   ./pokedex-cli
   ```

3. __Start exploring:__

   ```javascript
   Pokedex > help
   ```

You'll be greeted with the interactive Pokédex prompt where you can start your adventure!

## Usage

### Core Commands

#### Getting Help

```bash
Pokedex > help
```

Displays all available commands and their descriptions.

#### Location Exploration

```bash
# View next 20 locations
Pokedex > map

# View previous 20 locations  
Pokedex > mapb

# Explore a specific location
Pokedex > explore canalave-city-area
```

#### Pokémon Management

```bash
# Attempt to catch a Pokémon
Pokedex > catch pikachu

# Inspect a caught Pokémon's stats
Pokedex > inspect pikachu

# View your entire collection
Pokedex > pokedex
```

#### System Commands

```bash
# Exit the application
Pokedex > exit
```

### Example Session

```bash
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
# ... more locations

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - staryu
 - magikarp
 - gyarados

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
You may now inspect it with the inspect command!

Pokedex > inspect pikachu
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

Pokedex > pokedex
Your Pokedex:
        - pikachu
```

### Features

- __🗺️ Location Navigation__: Browse through Pokémon world locations with pagination
- __🔍 Area Exploration__: Discover which Pokémon inhabit specific areas
- __⚡ Pokémon Catching__: Attempt to catch Pokémon with realistic success rates
- __📊 Detailed Stats__: View comprehensive Pokémon information including stats and types
- __💾 Collection Management__: Build and view your personal Pokédex
- __🚀 Performance Caching__: Intelligent caching system for faster API responses
- __🎮 Interactive REPL__: Smooth command-line interface with helpful prompts

## Contributing

Contributions are welcome! This project demonstrates several Go programming concepts and is perfect for learning or extending.

### Development Setup

1. __Fork the repository__
2. __Clone your fork:__
   ```bash
   git clone https://github.com/your-username/pokedex-cli.git
   ```
3. __Create a feature branch:__
   ```bash
   git checkout -b feature/amazing-feature
   ```
4. __Make your changes and test:__
   ```bash
   go test ./...
   ```
5. __Commit and push:__
   ```bash
   git commit -m "Add amazing feature"
   git push origin feature/amazing-feature
   ```
6. __Open a Pull Request__

### Project Structure

```javascript
pokedex-cli/
├── main.go                     # Application entry point
├── cliHandler/                 # CLI command handling and logic
├── pokecache/                  # Caching system implementation
├── pokedexApiHandler/          # PokéAPI integration
└── README.md                   # This file
```

### Areas for Contribution

- __New Commands__: Add additional Pokédex functionality
- __UI Improvements__: Enhance the command-line interface
- __Performance__: Optimize caching strategies
- __Testing__: Expand test coverage
- __Documentation__: Improve code documentation

---

*Built with ❤️ and Go. Powered by [PokéAPI](https://pokeapi.co/)*
