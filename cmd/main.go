package main

import (
	"log"

	"github.com/hbjydev/zetman"
)

func main() {
	log.SetPrefix("[zetman] ")
	zetman.Cmd.Run()
}
