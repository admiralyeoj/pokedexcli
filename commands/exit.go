package commands

import (
	"os"

	"github.com/admiralyeoj/pokedexcli/types"
)

func Exit(cfg *types.Config) error {
	os.Exit(0)
	return nil
}
