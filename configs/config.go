package configs

import "github.com/admiralyeoj/pokedexcli/internal/pokeApi"

type Config struct {
	PokeApiClient    pokeApi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}
