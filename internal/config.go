package internal

import (
	"os"

	config "github.com/rwxrob/config/pkg"
)

func GetConfig() config.Conf {
	id := "zetman"
	dir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	return config.Conf{
		Id:   id,
		Dir:  dir,
		File: `config.yaml`,
	}
}
