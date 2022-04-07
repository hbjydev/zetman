package cmd

import Z "github.com/rwxrob/bonzai/z"

var TitlesCmd = &Z.Cmd{
	Name:    `titles`,
	Summary: `lists all zettelkasten and their titles`,

	Call: func(caller *Z.Cmd, args ...string) error {
		return nil
	},
}
