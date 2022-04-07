package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	Z "github.com/rwxrob/bonzai/z"
	config "github.com/rwxrob/config/pkg"
	"github.com/rwxrob/fs"
)

type configData struct {
	Path string
}

var InitCmd = &Z.Cmd{
	Name:    `init`,
	Summary: `initializes a zettelkasten`,

	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must specify a path")
		}

		rootDir, err := filepath.Abs(args[0])
		if err != nil {
			return err
		}

		isDir := fs.IsDir(rootDir)
		if isDir {
			return fmt.Errorf("%v: directory already exists", rootDir)
		}

		if err := os.Mkdir(rootDir, 0750); err != nil {
			return err
		}

		log.Printf("git init %v\n", rootDir)
		Z.Exec("git", "init", rootDir)
		log.Printf("gh repo create --source=%v --private\n", rootDir)
		Z.Exec("gh", "repo", "create", fmt.Sprintf("--source=%v", rootDir), "--private")

		newConfig := configData{
			Path: rootDir,
		}

		config.Init(x.Root.Name)
		config.Write(x.Root.Name, newConfig)

		return nil
	},
}
