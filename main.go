package main

import (
	"time"

	"github.com/admiralyeoj/pokedexcli/configs"
	"github.com/admiralyeoj/pokedexcli/internal/pokeApi"
)

func main() {
	pokeClient := pokeApi.NewClient(5*time.Second, time.Second*5)

	cfg := &configs.Config{
		PokeApiClient: pokeClient,
		CaughtPokemon: map[string]pokeApi.Pokemon{},
	}

	startRepl(cfg)
}
