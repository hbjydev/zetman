package main

import (
	"log"

	"github.com/hbjydev/zetman/internal"
)

func main() {
	log.SetPrefix("[zetman] ")
	internal.Cmd.Run()
}
