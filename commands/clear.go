package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/admiralyeoj/pokedexcli/configs"
)

func Clear(cfg *configs.Config) error {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J") // ANSI escape codes to clear screen
	}

	return nil
}
