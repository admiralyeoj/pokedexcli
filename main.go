package main

import (
	"time"

	"github.com/admiralyeoj/pokedexcli/internal/pokeApi"
	"github.com/admiralyeoj/pokedexcli/types"
)

func main() {
	pokeClient := pokeApi.NewClient(5 * time.Second)
	cfg := &types.Config{
		PokeApiClient: pokeClient,
	}

	startRepl(cfg)
}
