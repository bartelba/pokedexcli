package shared

import "github.com/bartelba/pokedexcli/internal/pokecache"

type Config struct {
    Next     string
    Previous string
    Cache    *pokecache.Cache
}
