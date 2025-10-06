Pokedex CLI
A command-line interface (CLI) written in Go that lets you catch, inspect, and list Pokémon using the PokeAPI. Built as part of the Boot.dev Go course.

🚀 Features
catch <pokemon> – Try to catch a Pokémon by name
inspect <pokemon> – View stats and details of a caught Pokémon
pokedex – List all Pokémon you've successfully caught
REPL interface with command parsing and error handling
Uses caching to reduce API calls

📦 Installation
bash
git clone https://github.com/your-username/pokedex-cli.git
cd pokedex-cli
go run .
Requires Go 1.20+ installed on your system.

🧪 Running the REPL
Once you've cloned the repo and have Go installed, you can start the REPL (Read-Eval-Print Loop) by running:

bash

go run .

This will launch an interactive prompt that looks like:

Code

Pokedex >

From here, you can enter commands like:

catch <pokemon> – Try to catch a Pokémon by name

inspect <pokemon> – View details of a caught Pokémon

pokedex – List all Pokémon you've caught

To exit the REPL, press Ctrl+C or type exit.

🕹️ Example Usage
bash
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
You may now inspect it with the inspect command.

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
  - hp: 35
  - attack: 55
  - defense: 40
  - special-attack: 50
  - special-defense: 50
  - speed: 90
Types:
  - electric

Pokedex > pokedex
Your Pokedex:
 - pikachu
🛠️ Tech Stack
Go (Golang)

PokeAPI

REPL loop

Custom command registry

📚 Possible Future Improvements
Save caught Pokémon to disk
Add evolution and leveling system
Simulate battles
Support different Pokéballs with varying catch rates

📜 License
MIT License
