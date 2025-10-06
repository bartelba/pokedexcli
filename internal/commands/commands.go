package commands

import (
    "fmt"
    "strings"
    "github.com/bartelba/pokedexcli/internal/shared"    
)

func Inspect(cfg *shared.Config, args ...string) error {
    if len(args) < 1 {
        fmt.Println("Usage: inspect <pokemon>")
        return nil
    }

    name := strings.ToLower(args[0])
    pokemon, ok := caughtPokemon[strings.ToLower(name)]
    if !ok {
        fmt.Println("you have not caught that pokemon")
        return nil
    }

    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)
    fmt.Println("Stats:")
    for stat, value := range pokemon.Stats {
        fmt.Printf("  -%s: %d\n", stat, value)
    }
    fmt.Println("Types:")
    for _, t := range pokemon.Types {
        fmt.Printf("  - %s\n", t)
    }

    return nil
}