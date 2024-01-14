package commands

import (
	"os"

	"github.com/admiralyeoj/pokedexcli/configs"
)

func Exit(cfg *configs.Config) error {
	os.Exit(0)
	return nil
}
