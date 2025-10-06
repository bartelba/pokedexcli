package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
    "os"
    "github.com/bartelba/pokedexcli/internal/commands"    
    "github.com/bartelba/pokedexcli/internal/shared"   
    "bufio"
)

func cleanInput(text string) []string {
    text = strings.TrimSpace(text)
    text = strings.ToLower(text)
    return strings.Fields(text)
}

type locationAreaResponse struct {
    Results  []struct {
        Name string `json:"name"`
    } `json:"results"`
    Next     string `json:"next"`
    Previous string `json:"previous"`
}


func commandMap(cfg *shared.Config, args ...string) error {
    url := cfg.Next
    if url == "" {
        url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
    }

    var body []byte
    if cached, ok := cfg.Cache.Get(url); ok {
        fmt.Println("Using cached data")
        body = cached
    } else {
        resp, err := http.Get(url)
        if err != nil {
            return err
        }
        defer resp.Body.Close()

        body, err = io.ReadAll(resp.Body)
        if err != nil {
            return err
        }

        cfg.Cache.Add(url, body)
    }

    var data locationAreaResponse
    if err := json.Unmarshal(body, &data); err != nil {
        return err
    }

    for _, loc := range data.Results {
        fmt.Println(loc.Name)
    }

    cfg.Next = data.Next
    cfg.Previous = data.Previous
    return nil
}

func commandMapb(cfg *shared.Config, args ...string) error {
    if cfg.Previous == "" {
        fmt.Println("you're on the first page")
        return nil
    }

    resp, err := http.Get(cfg.Previous)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var data locationAreaResponse
    if err := json.Unmarshal(body, &data); err != nil {
        return err
    }

    for _, loc := range data.Results {
        fmt.Println(loc.Name)
    }

    cfg.Next = data.Next
    cfg.Previous = data.Previous
    return nil
}

type cliCommand struct {
    name        string
    description string
    callback    func(*shared.Config, ...string) error
}

func commandExit(cfg *shared.Config, args ...string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil // unreachable, but satisfies the signature
}

func commandHelp(cfg *shared.Config, args ...string) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for _, cmd := range commandRegistry {
        fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
    }
    return nil
}

func commandExplore(cfg *shared.Config, args ...string) error {
    if len(args) < 1 {
        fmt.Println("Usage: explore <location-area>")
        return nil
    }

    area := args[0]
    url := "https://pokeapi.co/api/v2/location-area/" + area

    fmt.Printf("Exploring %s...\n", area)

    var body []byte
    if cached, ok := cfg.Cache.Get(url); ok {
        fmt.Println("Using cached data")
        body = cached
    } else {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Println("Error fetching location:", err)
            return nil
        }
        defer resp.Body.Close()

        body, err = io.ReadAll(resp.Body)
        if err != nil {
            fmt.Println("Error reading response:", err)
            return nil
        }

        cfg.Cache.Add(url, body)
    }

    var data struct {
        PokemonEncounters []struct {
            Pokemon struct {
                Name string `json:"name"`
            } `json:"pokemon"`
        } `json:"pokemon_encounters"`
    }

    if err := json.Unmarshal(body, &data); err != nil {
        fmt.Println("Error parsing JSON:", err)
        return nil
    }

    fmt.Println("Found Pokemon:")
    for _, encounter := range data.PokemonEncounters {
        fmt.Printf(" - %s\n", encounter.Pokemon.Name)
    }

    return nil
}

var commandRegistry map[string]cliCommand

func init() {
    commandRegistry = map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Explore the next 20 location areas",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Go back to the previous 20 location areas",
            callback:    commandMapb,
        },
        "explore": {
            name:        "explore",
            description: "Explore a specific location area",
            callback:    commandExplore,
        },
        "catch": {
            name:        "catch",
            description: "Catch a Pokémon by name",
            callback: func(cfg *shared.Config, args ...string) error {
                if len(args) < 1 {
                    fmt.Println("Usage: catch <pokemon>")
                    return nil
                }
                return commands.Catch(cfg, args...)
            },
        },        
        "inspect": {
            name:        "inspect",
            description: "Inspect a caught Pokémon by name",
            callback: func(cfg *shared.Config, args ...string) error {
                if len(args) < 1 {
                    fmt.Println("Usage: inspect <pokemon>")
                    return nil
                }
                commands.Inspect(cfg, args...)
                return nil
            },
        },        
        "pokedex": {
            name:        "pokedex",
            description: "List all caught Pokémon",
            callback: func(cfg *shared.Config, args ...string) error {
                return commands.Pokedex(cfg, args...)
             },
        },
    }
}

func startRepl(cfg *shared.Config) {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        userInput := scanner.Text()

        input := cleanInput(userInput)
        if len(input) == 0 {
            continue
        }

        cmdName := input[0]
        args := input[1:]

        if cmd, ok := commandRegistry[cmdName]; ok {
            cmd.callback(cfg, args...)
        } else {
            fmt.Println("Unknown command:", cmdName)
        }
    }
}
