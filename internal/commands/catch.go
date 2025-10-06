package commands

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "strings"
    "time"
	"github.com/bartelba/pokedexcli/internal/shared"
)

type Pokemon struct {
    Name           string            `json:"name"`
    BaseExperience int               `json:"base_experience"`
    Height         int               `json:"height"`
    Weight         int               `json:"weight"`
    Stats          map[string]int    `json:"stats"`
    Types          []string          `json:"types"`
}

// var pokedex = make(map[string]Pokemon)
var caughtPokemon = make(map[string]Pokemon)

func Catch(cfg *shared.Config, args ...string) error {
    if len(args) < 1 {
        fmt.Println("Usage: catch <pokemon-name>")
        return nil
    }

    name := strings.ToLower(args[0])
    fmt.Printf("Throwing a Pokeball at %s...\n", name)

    pokemon, err := fetchPokemon(name)
    if err != nil {
        fmt.Println("Could not find that Pokémon.")
        return nil
    }

    rand.Seed(time.Now().UnixNano())
    chance := 100 - pokemon.BaseExperience
    if chance < 10 {
        chance = 10
    }

    if rand.Intn(100) < chance {
        fmt.Printf("%s was caught!\n", pokemon.Name)
        fmt.Println("You may now inspect it with the inspect command.")
        caughtPokemon[strings.ToLower(pokemon.Name)] = pokemon
    } else {
        fmt.Printf("%s escaped!\n", pokemon.Name)
    }

    return nil
}

func fetchPokemon(name string) (Pokemon, error) {
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
    resp, err := http.Get(url)
    if err != nil || resp.StatusCode != 200 {
        return Pokemon{}, fmt.Errorf("failed to fetch Pokémon")
    }
    defer resp.Body.Close()

    var raw struct {
        Name           string `json:"name"`
        BaseExperience int    `json:"base_experience"`
        Height         int    `json:"height"`
        Weight         int    `json:"weight"`
        Stats          []struct {
            Stat struct {
                Name string `json:"name"`
            } `json:"stat"`
            BaseStat int `json:"base_stat"`
        } `json:"stats"`
        Types []struct {
            Type struct {
                Name string `json:"name"`
            } `json:"type"`
        } `json:"types"`
    }

    err = json.NewDecoder(resp.Body).Decode(&raw)
    if err != nil {
        return Pokemon{}, err
    }

    stats := make(map[string]int)
    for _, s := range raw.Stats {
        stats[s.Stat.Name] = s.BaseStat
    }

    types := make([]string, 0)
    for _, t := range raw.Types {
        types = append(types, t.Type.Name)
    }

    return Pokemon{
        Name:           raw.Name,
        BaseExperience: raw.BaseExperience,
        Height:         raw.Height,
        Weight:         raw.Weight,
        Stats:          stats,
        Types:          types,
    }, nil
}