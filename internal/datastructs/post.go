package datastructs

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hbjydev/zetman/internal/consts"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/fs"
	"github.com/rwxrob/fs/file"
)

// Post represents a post in the Zettelkasten
type Post struct {
	// Id is the ID of the Post, which is an isosec.
	Id string `json:"id" yaml:"id"`

	// The path to the Post
	Path string `json:"path" yaml:"path"`
}

// Touch will create and edit the file at the path specified by the Post.
func (p *Post) Touch() error {
	if exists := fs.IsDir(p.Path); exists {
		return fmt.Errorf("dir already exists")
	}

	if err := os.Mkdir(p.Path, 0750); err != nil {
		return err
	}

	mdPath := filepath.Join(p.Path, consts.DefaultName)

	if exists := fs.Exists(mdPath); exists {
		return fmt.Errorf("file already exists")
	}

	err := file.Edit(mdPath)
	if err != nil {
		return err
	}

	if err := os.Chdir(p.Path); err != nil {
		return err
	}
	if err := Z.Exec("git", "add", "."); err != nil {
		return err
	}

	if err := Z.Exec("git", "commit", "-m", fmt.Sprintf("Added post %v", p.Id)); err != nil {
		return err
	}

	branch, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return err
	}
	branchStr := strings.Trim(string(branch), "\n")

	if err := Z.Exec("git", "push", "origin", branchStr); err != nil {
		return err
	}

	return nil
}

// Edit opens up the user's default editor on the README.md file for the Post.
func (p *Post) Edit() error {
	if exists := fs.IsDir(p.Path); !exists {
		return fmt.Errorf("dir does not exist")
	}

	mdPath := filepath.Join(p.Path, consts.DefaultName)

	if exists := fs.Exists(mdPath); !exists {
		return fmt.Errorf("file does not exist")
	}

	return file.Edit(mdPath)
}
