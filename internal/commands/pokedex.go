package commands

import (
    "fmt"
    "sort"
    "github.com/bartelba/pokedexcli/internal/shared"    	
)

func Pokedex(cfg *shared.Config, args ...string) error {
    if len(caughtPokemon) == 0 {
        fmt.Println("You haven't caught any Pokémon yet.")
        return nil
    }

    fmt.Println("Your Pokedex:")
    names := make([]string, 0, len(caughtPokemon))
    for name := range caughtPokemon {
        names = append(names, name)
    }

    sort.Strings(names) // Optional: alphabetize
    for _, name := range names {
        fmt.Printf(" - %s\n", name)
    }

    return nil
}
