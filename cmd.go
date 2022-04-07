package zetman

import (
	cmd "github.com/hbjydev/zetman/internal/cmd"
	"github.com/rwxrob/bonzai/help"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/config"
)

var Cmd = &Z.Cmd{
	Name:      `zetman`,
	Summary:   `my zettelkasten management tool`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 Hayden Young`,
	License:   `apache-2.0`,
	Commands:  []*Z.Cmd{help.Cmd, config.Cmd, cmd.InitCmd, cmd.PostCmd, cmd.TitlesCmd},
}
