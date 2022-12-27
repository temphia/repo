package code

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/temphia/temphia/code/core/backend/libx/xutils"
)

func (rb *RepoBuild) buildItem(name string) (string, error) {

	fmt.Println("Building ", name)

	item := rb.Config.Items[name]

	buildPath := path.Join(rb.Config.BuildFolder, name)
	outputPath := path.Join(rb.Config.OutputFolder, name, item.OutputFolder)

	err := xutils.CreateIfNotExits(buildPath)
	if err != nil {
		return "", err
	}

	_, err = git.PlainClone(buildPath, false, &git.CloneOptions{
		URL:           item.GitURL,
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewBranchReferenceName(item.Branch),
		SingleBranch:  true,
		Depth:         1,
	})

	if err != nil {
		if !errors.Is(git.ErrRepositoryAlreadyExists, err) {
			return "", err
		}
	}

	curr, _ := os.Getwd()

	cmd := exec.Command(item.BuildCommand)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = path.Join(curr, buildPath)

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return outputPath, CopyDirectory(buildPath, outputPath)
}
