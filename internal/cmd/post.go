package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/hbjydev/zetman/internal"
	"github.com/hbjydev/zetman/internal/datastructs"
	Z "github.com/rwxrob/bonzai/z"
	uniq "github.com/rwxrob/uniq-go"
)

var PostCmd = &Z.Cmd{
	Name:    `post`,
	Summary: `post a new slip into the zettelkasten`,

	Call: func(caller *Z.Cmd, args ...string) error {
		conf := internal.GetConfig()

		path := conf.Query(".path")
		if path == "" {
			return fmt.Errorf("zettelkasten not initialized, run zetman init")
		}

		id := uniq.IsoSecond()
		post := datastructs.Post{
			Id:   id,
			Path: filepath.Join(path, id),
		}

		return post.Touch()
	},
}
